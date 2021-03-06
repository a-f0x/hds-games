package rabbit

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

/*
Умеет в рекконект консюмеров только к очередям с роутинг кеем!
*/
type AmqpConsumer struct {
	client    *amqpClient
	startChan chan bool
	streams   map[string]*stream
}

func NewAmqpConsumer(host string, port int64, user string, password string, reconnectionTimeSec int32) *AmqpConsumer {
	c := AmqpConsumer{
		client:  newAmqpClient(host, port, user, password, reconnectionTimeSec),
		streams: make(map[string]*stream),
	}
	connectionInfo := c.client.connect()
	c.startChan = make(chan bool)
	go c.watch(connectionInfo)
	return &c
}

func (ac *AmqpConsumer) reconnect() {
	for _, stream := range ac.streams {
		ac.connectStream(stream)
	}
}

func (ac *AmqpConsumer) streaming() {
	for _, stream := range ac.streams {
		if stream.incomingChannel == nil {
			ac.connectStream(stream)
		}
	}
}

func (ac *AmqpConsumer) connectStream(stream *stream) {
	log.Printf("connect stream %s", stream.queue)
	messages, err := ac.client.channel.Consume(
		stream.queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Queue %s consumer error %s", stream.queue, err)
	}
	stream.incomingChannel = messages
	stream.consume()
}

//как тут сделать отписку и отмену ? как то видимо с контекстом работать надо. я еще пока до него не добрался.
func (ac *AmqpConsumer) watch(connectionInfo <-chan bool) {
	for {
		select {
		case cInfo := <-connectionInfo:
			//log.Printf("consumer connection state %v", cInfo)
			if cInfo {
				ac.reconnect()
			}
		case startStreaming := <-ac.startChan:
			//log.Printf("consumer streaming state %v", startStreaming)
			if startStreaming && ac.client.isConnected() {
				ac.streaming()
			}
		}
	}

}

func (ac *AmqpConsumer) Subscribe(ctx context.Context, queue string) (<-chan []byte, error) {
	alreadyExistStream := ac.streams[queue]
	if alreadyExistStream != nil {
		return nil, errors.New(fmt.Sprintf("already subscribed to queue %s", queue))
	}
	stream := &stream{
		ctx:             ctx,
		queue:           queue,
		outChan:         make(chan []byte),
		incomingChannel: nil,
	}
	ac.streams[queue] = stream
	ac.startChan <- true
	return stream.outChan, nil
}
