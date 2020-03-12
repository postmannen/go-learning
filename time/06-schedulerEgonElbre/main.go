package main

import (
	"container/heap"
	"context"
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	ctx context.Context

	mu       sync.Mutex
	jobadded chan struct{}
	items    tasks
}

type tasks []task

type task struct {
	Deadline time.Time
	Call     func()
}

func NewQueue(ctx context.Context) *Queue {
	return &Queue{ctx: ctx, jobadded: make(chan struct{})}
}

func (q *Queue) monitor() {
	timer := time.NewTimer(0)
	defer timer.Stop()
	for {
		now := time.Now()
		var tick <-chan time.Time

		next := q.flush(now)
		if !next.IsZero() {
			timer.Reset(next.Sub(now))
			tick = timer.C
		} else {
			timer.Stop()
		}

		select {
		case <-q.ctx.Done():
			return
		case <-q.jobadded:
		case <-tick:
		}
	}
}

func (q *Queue) flush(now time.Time) (next time.Time) {
	q.mu.Lock()
	defer q.mu.Unlock()

	for q.items.Len() > 0 {
		next := &q.items[0]
		if next.Deadline.After(now) {
			return next.Deadline
		}
		item := heap.Pop(&q.items).(task)
		go item.Call()
	}

	return time.Time{}
}

func (q *Queue) Delay(duration time.Duration, fn func()) {
	q.mu.Lock()
	defer q.mu.Unlock()

	heap.Push(&q.items, task{
		Deadline: time.Now().Add(duration),
		Call:     fn,
	})

	select {
	case q.jobadded <- struct{}{}:
	default:
	}
}

func main() {
	q := NewQueue(context.Background())
	go q.monitor()

	q.Delay(5*time.Second, func() { fmt.Println("5") })
	q.Delay(2*time.Second, func() { fmt.Println("2") })
	q.Delay(3*time.Second, func() { fmt.Println("3") })
	q.Delay(1*time.Second, func() { fmt.Println("1") })
	q.Delay(4*time.Second, func() { fmt.Println("4") })

	time.Sleep(6 * time.Second)
}

func (pq tasks) Len() int            { return len(pq) }
func (pq tasks) Less(i, j int) bool  { return pq[i].Deadline.Before(pq[j].Deadline) }
func (pq tasks) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *tasks) Push(x interface{}) { *pq = append(*pq, x.(task)) }
func (pq *tasks) Pop() (item interface{}) {
	n := len(*pq)
	item, *pq = (*pq)[n-1], (*pq)[:n-1]
	return item
}

