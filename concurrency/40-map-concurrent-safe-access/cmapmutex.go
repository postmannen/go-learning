package main

import (
	"sync"
)

type cMapMutex struct {
	m  map[int]string
	mu sync.Mutex
}

func newCMapMutex() *cMapMutex {
	cM := cMapMutex{
		m: map[int]string{},
	}
	return &cM
}

func (c *cMapMutex) put(kv keyValue) {
	c.mu.Lock()
	c.m[kv.k] = kv.v
	c.mu.Unlock()
}

func (c *cMapMutex) get(key int) keyValue {
	var kv keyValue

	c.mu.Lock()
	kv.v, kv.ok = c.m[key]
	c.mu.Unlock()

	return kv
}

func (c *cMapMutex) del(kv keyValue) {
	c.mu.Lock()
	delete(c.m, kv.k)
	c.mu.Unlock()

}

func (c *cMapMutex) getAll() []keyValue {
	var kvs []keyValue

	c.mu.Lock()
	for k, v := range c.m {
		kvs = append(kvs, keyValue{k: k, v: v})
	}
	c.mu.Unlock()

	return kvs
}
