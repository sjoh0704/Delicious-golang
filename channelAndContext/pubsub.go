package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// publisher
type Publisher struct {
	ctx         context.Context
	subscribeCh chan chan<- string // channel을 chan<-string 타입으로 만들겠다.  스트링 데이터를 넣기만 가능한 채널 
	publishCh   chan string
	subscribers []chan<- string
}

func NewPublisher(ctx context.Context) *Publisher {
	return &Publisher{
		ctx:         ctx,
		subscribeCh: make(chan chan<- string),
		publishCh:   make(chan string),
		subscribers: make([]chan<- string, 0),
	}
}

func (p *Publisher) Subscribe(sub chan<- string) {
	p.subscribeCh <- sub
}

func (p *Publisher) Publish(msg string) {
	p.publishCh <- msg
}

func (p *Publisher) Update() {
	for {
		select {
		case sub := <-p.subscribeCh:
			p.subscribers = append(p.subscribers, sub)
		case msg := <-p.publishCh:
			for _, subscriber := range p.subscribers {
				subscriber <- msg
			}
		case <-p.ctx.Done():
			wg.Done()
			return
		}
	}
}

// subscriber

type Subscriber struct {
	ctx   context.Context
	name  string
	msgCh chan string
}

func NewSubscriber(name string, ctx context.Context) *Subscriber {
	return &Subscriber{
		ctx:   ctx,
		name:  name,
		msgCh: make(chan string),
	}
}

func (s *Subscriber) Subscribe(pub *Publisher) {
	pub.Subscribe(s.msgCh)
}

func (s *Subscriber) Update() {
	for {
		select {
		case msg := <-s.msgCh:
			fmt.Printf("%s got Message:%s\n", s.name, msg)
		case <-s.ctx.Done():
			wg.Done()
			return
		}
	}
}

// main


func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(4)
	publisher := NewPublisher(ctx)
	subscriber1 := NewSubscriber("AAA", ctx)
	subscriber2 := NewSubscriber("BBB", ctx)

	go publisher.Update()

	subscriber1.Subscribe(publisher)
	subscriber2.Subscribe(publisher)

	go subscriber1.Update()
	go subscriber2.Update()

	go func() {
		tick := time.Tick(time.Second * 2)
		for {
			select {
			case <-tick:
				publisher.Publish("Hello Message")
			case <-ctx.Done():
				wg.Done()
				return
			}
		}
	}()

	fmt.Scanln() // 키보드 입력을 기다리므로 여기서 멈춘다. 
	cancel()

	wg.Wait()
}