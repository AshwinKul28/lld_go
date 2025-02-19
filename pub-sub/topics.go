package main

import "sync"

type Topic struct {
	Name        string
	Subscribers []chan string
	mu          sync.Mutex
}

func NewTopic(name string) *Topic {
	return &Topic{
		Name: name,
	}
}

func (t *Topic) Subscribe(sub *Subscriber) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Subscribers = append(t.Subscribers, sub.Ch)
}

func (t *Topic) Publish(message string) {
	for _, sub := range t.Subscribers {
		sub <- message
	}
}
