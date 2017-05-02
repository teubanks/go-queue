package queue

import "sync"

type Queue struct {
	nodes []interface{}
	head  int
	tail  int
	cnt   int
}

var QueueMutex sync.RWMutex

func NewQueue() *Queue {
	return &Queue{
		nodes: make([]interface{}, 2),
	}
}

func (q *Queue) resize(n int) {
	nodes := make([]interface{}, n)
	if q.head < q.tail {
		copy(nodes, q.nodes[q.head:q.tail])
	} else {
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.tail])
	}

	q.tail = q.cnt % n
	q.head = 0
	q.nodes = nodes
}

// Push stores an object into the queue. If the queue gets pushed beyond its
// current bounds, it's size gets doubled
func (q *Queue) Push(i interface{}) {
	QueueMutex.Lock()
	defer QueueMutex.Unlock()

	if q.cnt == len(q.nodes) {
		// Also tested a grow rate of 1.5, see: http://stackoverflow.com/questions/2269063/buffer-growth-strategy
		// In Go this resulted in a higher memory usage.
		q.resize(q.cnt * 2)
	}
	q.nodes[q.tail] = i
	q.tail = (q.tail + 1) % len(q.nodes)
	q.cnt++
}

// Pop removes the first object from the queue. If number of elements in the
// queue goes below half of available space, the queue capacity gets halved
func (q *Queue) Pop() (interface{}, bool) {
	QueueMutex.Lock()
	defer QueueMutex.Unlock()

	if q.cnt == 0 {
		return nil, false
	}
	i := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.cnt--

	if n := len(q.nodes) / 2; n > 2 && q.cnt <= n {
		q.resize(n)
	}

	return i, true
}

// Cap returns the returns the current capacity of the queue
func (q *Queue) Cap() int {
	QueueMutex.RLock()
	defer QueueMutex.RUnlock()

	return cap(q.nodes)
}

// Len returns the number of queue objects currently stored in the queue
func (q *Queue) Len() int {
	QueueMutex.RLock()
	defer QueueMutex.RUnlock()

	return q.cnt
}
