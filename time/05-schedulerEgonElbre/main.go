package main

import (
	"container/heap"
	"context"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

// A Group is a goroutine worker pool which schedules tasks to be performed
// after a specified time. A Group must be created with the New constructor.
type Group struct {
	// Task runner and a heap of tasks to be run.
	running *errgroup.Group
	mu      sync.Mutex
	tasks   tasks

	// Signals for when a task is added and how many tasks remain on the heap.
	added chan struct{}
}

// New creates a new Group which will use ctx for cancelation. If cancelation
// is not a concern, use context.Background().
func New() *Group {
	return &Group{
		running: &errgroup.Group{},
		added:   make(chan struct{}),
	}
}

// Delay schedules a function to run at or after the specified delay. Delay
// is a convenience wrapper for Schedule which adds delay to the current time.
// Specifying a negative delay will cause the task to be scheduled immediately.
func (g *Group) Delay(delay time.Duration, fn func() error) {
	g.Schedule(time.Now().Add(delay), fn)
}

// Schedule schedules a function to run at or after the specified time.
// Specifying a past time will cause the task to be scheduled immediately.
func (g *Group) Schedule(when time.Time, fn func() error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	heap.Push(&g.tasks, task{
		Deadline: when,
		Call:     fn,
	})

	// Notify monitor that a new task has been pushed on to the heap.
	select {
	case g.added <- struct{}{}:
	default:
	}
}

// Run will run the group.
func (g *Group) Run(ctx context.Context) error {
	t := time.NewTimer(0)
	defer t.Stop()

	for {
		if ctx.Err() != nil {
			// Context canceled.
			return g.running.Wait()
		}

		now := time.Now()
		var tick <-chan time.Time

		// Start any tasks that are ready as of now.
		next := g.trigger(now)
		if !next.IsZero() {
			// Wait until the next scheduled task is ready.
			t.Reset(next.Sub(now))
			tick = t.C
		} else {
			t.Stop()
		}

		select {
		case <-ctx.Done():
			// Context canceled.
			return g.running.Wait()
		case <-g.added:
			// A new task was added, check task heap again.
		case <-tick:
			// An existing task should be ready as of now.
		}
	}
}

// trigger checks for scheduled tasks and runs them if they are scheduled
// on or after the time specified by now.
func (g *Group) trigger(now time.Time) time.Time {
	g.mu.Lock()
	defer g.mu.Unlock()

	for g.tasks.Len() > 0 {
		next := &g.tasks[0]
		if next.Deadline.After(now) {
			// Earliest scheduled task is not ready.
			return next.Deadline
		}

		// This task is ready, pop it from the heap and run it.
		t := heap.Pop(&g.tasks).(task)
		g.running.Go(t.Call)
	}

	return time.Time{}
}

// A task is a function which is called after the specified deadline.
type task struct {
	Deadline time.Time
	Call     func() error
}

// tasks implements heap.Interface.
type tasks []task

var _ heap.Interface = &tasks{}

func (pq tasks) Len() int            { return len(pq) }
func (pq tasks) Less(i, j int) bool  { return pq[i].Deadline.Before(pq[j].Deadline) }
func (pq tasks) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *tasks) Push(x interface{}) { *pq = append(*pq, x.(task)) }
func (pq *tasks) Pop() (item interface{}) {
	n := len(*pq)
	item, *pq = (*pq)[n-1], (*pq)[:n-1]
	return item
}
