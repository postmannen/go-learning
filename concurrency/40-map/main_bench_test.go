package main

import (
	"context"
	"sync"
	"testing"
)

var result string

func BenchmarkPut(b *testing.B) {
	cM := newCMap()
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		// fmt.Printf("Canceling context\n")
		cancel()
	}()

	go func() {
		cM.run(ctx)
	}()

	for i := 0; i < b.N; i++ {
		cM.put(keyValue{k: i, v: "a"})
	}

}

func BenchmarkPutNormal(b *testing.B) {
	cM := newCMap()

	var mu sync.Mutex

	for i := 0; i < b.N; i++ {
		mu.Lock()
		cM.m[i] = "a"
		mu.Unlock()
	}
}

var resultKv keyValue

func BenchmarkGet(b *testing.B) {
	cM := newCMap()
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		// fmt.Printf("Canceling context\n")
		cancel()
	}()

	go func() {
		cM.run(ctx)
	}()

	cM.put(keyValue{k: 1, v: "a"})

	var kv keyValue
	for i := 0; i < b.N; i++ {
		kv = cM.get(1)

	}
	resultKv = kv

}

func BenchmarkGetNormal(b *testing.B) {
	cM := newCMap()

	var mu sync.Mutex

	cM.m[1] = "a"
	var v string

	for i := 0; i < b.N; i++ {
		mu.Lock()
		v = cM.m[1]
		mu.Unlock()
	}

	result = v
}

func BenchmarkGetWhileConcurrentPut(b *testing.B) {
	cM := newCMap()
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		// fmt.Printf("Canceling context\n")
		cancel()
	}()

	go func() {
		cM.run(ctx)
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				cM.put(keyValue{k: 1, v: "a"})
			}
		}
	}()

	var kv keyValue
	for i := 0; i < b.N; i++ {
		kv = cM.get(1)
	}
	resultKv = kv

}

func BenchmarkGetNormalWhileConcurrentPut(b *testing.B) {
	cM := newCMap()
	ctx, cancel := context.WithCancel(context.Background())
	var mu sync.Mutex

	defer func() {
		// fmt.Printf("Canceling context\n")
		cancel()
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				mu.Lock()
				cM.m[1] = "a"
				mu.Unlock()
			}
		}
	}()

	var v string
	for i := 0; i < b.N; i++ {
		mu.Lock()
		v = cM.m[1]
		mu.Unlock()
	}
	result = v
}
