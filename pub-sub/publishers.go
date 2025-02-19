package main

import "fmt"

type Publisher struct {
	ID     int
	System *PubSub
}

func NewPublisher(sys *PubSub) *Publisher {
	return &Publisher{
		System: sys,
	}
}

func (p *Publisher) Publish(topic string, message string) {
	if _, ok := p.System.Topics[topic]; !ok {
		fmt.Println("cannot send to non existing topic!")
		return
	}

	p.System.Topics[topic].Publish(message)
}
