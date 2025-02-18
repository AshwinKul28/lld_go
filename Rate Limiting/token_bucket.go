package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	Capacity  int
	Rate      time.Duration
	Tokens    chan struct{}
	CloseChan chan struct{}
}

func NewTokenBucket(capacity int, rate time.Duration) *TokenBucket {
	tb := &TokenBucket{
		Capacity:  capacity,
		Rate:      rate,
		Tokens:    make(chan struct{}, capacity),
		CloseChan: make(chan struct{}),
	}

	for i := 0; i < capacity; i++ {
		tb.Tokens <- struct{}{}
	}

	return tb
}

func (tb *TokenBucket) refillToken() {
	ticket := time.NewTicker(tb.Rate)
	defer ticket.Stop()

	for {
		select {
		case <-ticket.C:
			if len(tb.Tokens) < tb.Capacity {
				tb.Tokens <- struct{}{}
			}
		case <-tb.CloseChan:
			return
		}
	}
}

func (tb *TokenBucket) Allow() bool {
	select {
	case <-tb.Tokens:
		return true // Allowed, token consumed
	default:
		return false // Rate limit exceeded
	}
}

func (tb *TokenBucket) Stop() {
	close(tb.CloseChan)
}

func main() {
	tb := NewTokenBucket(5, 500*time.Millisecond) // Bucket size 5, refill every 500ms
	defer tb.Stop()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if tb.Allow() {
				fmt.Printf("Request %d allowed\n", id)
			} else {
				fmt.Printf("Request %d denied (Rate limited)\n", id)
			}
		}(i)
		time.Sleep(500 * time.Millisecond) // Simulate incoming requests
	}

	wg.Wait()
}
