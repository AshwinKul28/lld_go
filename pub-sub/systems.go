package main

import (
	"fmt"
	"sync"
	"time"
)

type PubSub struct {
	Topics map[string]*Topic
	mu     sync.RWMutex
}

func NewPubSub() *PubSub {
	return &PubSub{
		Topics: make(map[string]*Topic),
	}
}

func (s *PubSub) RegisterTopic(t *Topic) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Topics[t.Name] = t
}

func main() {
	system := NewPubSub()

	topic1 := NewTopic("INEVENTS")
	topic2 := NewTopic("IMAGES")
	topic3 := NewTopic("NEWS")

	system.RegisterTopic(topic1)
	system.RegisterTopic(topic2)
	system.RegisterTopic(topic3)

	ch1 := make(chan string)
	ch2 := make(chan string)
	sub1 := NewSubscriber(ch1)
	sub2 := NewSubscriber(ch2)

	topic1.Subscribe(sub1)
	topic1.Subscribe(sub2)
	topic2.Subscribe(sub2)
	topic3.Subscribe(sub1)

	pub1 := NewPublisher(system)
	pub2 := NewPublisher(system)

	go func() {
		pub1.Publish("INEVENTS", "hello from pub1 to events")
		pub1.Publish("NEWS", "hello from pub1 to bbc news")

		pub2.Publish("NEWS", "hello from pub2 to abc news")
		pub2.Publish("IMAGES", "hello from pub2 to images topic")
	}()

	for {
		select {
		case msg := <-ch1:
			fmt.Printf("Subscriber 1 received: %s\n", msg)
		case msg := <-ch2:
			fmt.Printf("Subscriber 2 received: %s\n", msg)
		case <-time.After(time.Second):
			// Add a small delay to prevent CPU spinning
			continue
		}
	}

}
