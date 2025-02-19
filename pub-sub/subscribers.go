package main

type Subscriber struct {
	ID int
	Ch chan string
}

func NewSubscriber(ch chan string) *Subscriber {
	return &Subscriber{
		Ch: ch,
	}
}
