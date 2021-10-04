package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
)

type keyValue struct {
	k  string
	v  string
	ok bool
}

type kvCh chan keyValue

type getValue struct {
	k    string
	kvCh kvCh
}

type cMap struct {
	m      map[string]string
	mInCh  chan kvCh
	mGetCh chan getValue
	mDelCh chan kvCh
}

func newCMap() *cMap {
	cM := cMap{
		m:      map[string]string{},
		mInCh:  make(chan kvCh),
		mGetCh: make(chan getValue),
		mDelCh: make(chan kvCh),
	}
	return &cM
}

func (c *cMap) run(ctx context.Context) {
	for {
		select {
		case kvCh := <-c.mInCh:
			kv := <-kvCh
			c.m[kv.k] = kv.v

		case gv := <-c.mGetCh:
			v, ok := c.m[gv.k]
			gv.kvCh <- keyValue{gv.k, v, ok}

		case kvCh := <-c.mDelCh:
			kv := <-kvCh
			delete(c.m, kv.k)

		case <-ctx.Done():
			log.Printf("info: cMap: got ctx.Done\n")
			return
		}
	}
}

func (c *cMap) put(kv keyValue) {
	kvCh := make(chan keyValue, 1)
	kvCh <- kv
	c.mInCh <- kvCh

}

func (c *cMap) get(key string) keyValue {
	kvCh := make(chan keyValue, 1)

	gv := getValue{
		k:    key,
		kvCh: kvCh,
	}

	c.mGetCh <- gv

	return <-kvCh
}

func (c *cMap) del(kv keyValue) {
	kvCh := make(chan keyValue, 1)
	kvCh <- kv
	c.mDelCh <- kvCh

}

func main() {
	var wg sync.WaitGroup
	cM := newCMap()
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		cM.run(ctx)
		wg.Done()
	}()

	// Fill and read concurrently.
	{
		var wg sync.WaitGroup
		const times = 3

		for i := 0; i < times; i++ {
			// Fill values
			wg.Add(1)
			go func() {
				for ii := 0; ii < 100; ii++ {
					s := strconv.Itoa(ii)
					sTen := strconv.Itoa(ii * 10)
					cM.put(keyValue{k: s, v: sTen})
				}
				wg.Done()
			}()

			wg.Add(1)
			go func() {
				kv := cM.get("3")
				if !kv.ok {
					fmt.Printf("info: no value found when key = %v\n", kv.k)
					return
				}

				fmt.Printf("key: %v, value: %v\n", kv.k, kv.v)
				wg.Done()
			}()
		}

		wg.Wait()
	}

	{
		// Deletion of key
		cM.del(keyValue{k: "50"})
		kv := cM.get("50")
		if !kv.ok {
			fmt.Printf("info: no value found when key = %v\n", kv.k)
			return
		}

		fmt.Printf("key: %v, value: %v\n", kv.k, kv.v)
	}

	cancel()
	wg.Wait()

}
