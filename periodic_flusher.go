package queue

import (
	"log"
	"time"
)

type PeriodicFlusher struct {
	queue        *Queue
	ticker       <-chan time.Time
	doneFlushing chan bool

	PostFlush func([]Queueable)
}

// TODO: Pass in a stream here
func NewPeriodicFlusher() *PeriodicFlusher {
	r := new(PeriodicFlusher)
	r.queue = NewQueue()
	r.ticker = time.NewTicker(3 * time.Second).C // publish results every 3 seconds
	r.doneFlushing = make(chan bool)

	return r
}

func (r *PeriodicFlusher) Enqueue(obj Queueable) {
	r.queue.Push(obj)
}

func (r *PeriodicFlusher) startFlusher() {
	r.flushQueue()

	for {
		log.Printf("flushing")
		select {
		case <-r.ticker:
			r.flushQueue()
		case <-r.doneFlushing:
			r.flushQueue()
			return
		}
	}
}

func (r *PeriodicFlusher) Start() {
	log.Println("starting queue")
	go r.startFlusher()
}

func (r *PeriodicFlusher) Stop() {
	log.Println("stopping queue")
	close(r.doneFlushing) // stops goroutine that flushes the pool
	// Flush one more time just in case the ticker
	// was stopped before all metrics were flushed
	r.flushQueue()
}

func (r *PeriodicFlusher) flushQueue() {
	var objects []Queueable

	for {
		m, ok := r.queue.Pop()
		if !ok {
			break
		}

		objects = append(objects, m)
	}

	r.PostFlush(objects)
}
