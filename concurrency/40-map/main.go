package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"
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
	m         map[string]string
	mInCh     chan kvCh
	mGetCh    chan getValue
	mDelCh    chan kvCh
	mGetAllCh chan chan []keyValue
}

func newCMap() *cMap {
	cM := cMap{
		m:         map[string]string{},
		mInCh:     make(chan kvCh),
		mGetCh:    make(chan getValue),
		mDelCh:    make(chan kvCh),
		mGetAllCh: make(chan chan []keyValue),
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

		case gaCh := <-c.mGetAllCh:
			kvSlice := []keyValue{}

			for k, v := range c.m {
				kv := keyValue{k: k, v: v}
				kvSlice = append(kvSlice, kv)
			}

			gaCh <- kvSlice

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

func (c *cMap) getAll() []keyValue {
	gaCh := make(chan []keyValue, 1)
	c.mGetAllCh <- gaCh

	all := <-gaCh
	return all
}

func main() {
	runtime.GOMAXPROCS(1)

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
		const keysValues = 20

		fmt.Println("-----------------------concurrently put and get--------------------------------------")

		for i := 0; i < times; i++ {
			// Fill values
			wg.Add(1)
			go func() {
				for ii := 0; ii < keysValues; ii++ {
					s := strconv.Itoa(ii)
					sTen := strconv.Itoa(ii * 10)
					cM.put(keyValue{k: s, v: sTen})
				}
				wg.Done()
			}()

			wg.Add(1)
			go func() {
				defer wg.Done()
				time.Sleep(time.Millisecond * 100)
				kv := cM.get("5")
				if !kv.ok {
					fmt.Printf("info: no value found when key = %v\n", kv.k)
					return
				}

				fmt.Printf("key: %v, value: %v\n", kv.k, kv.v)
			}()
		}

		wg.Wait()
	}

	{
		// Deletion of key
		fmt.Println("-----------------------key deletion--------------------------------------")
		fmt.Println("deleting key")
		cM.del(keyValue{k: "5"})
		kv := cM.get("5")
		if !kv.ok {
			fmt.Printf("info: delete, no value found when key = %v\n", kv.k)
		} else {
			fmt.Printf("key: %v, value: %v\n", kv.k, kv.v)
		}

	}

	{
		// Get all key/values
		fmt.Println("-----------------------get all values--------------------------------------")
		kvAll := cM.getAll()
		fmt.Printf("Len of slice with getAll : %v\n", len(kvAll))
		fmt.Printf("All values returned as a slice:\n %#v\n", kvAll)
	}

	cancel()
	wg.Wait()

}
