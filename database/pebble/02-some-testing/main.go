// Copyright 2020 The LevelDB-Go and Pebble Authors. All rights reserved. Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
)

type data struct {
	key   string
	value string
}

func main() {
	d := []data{data{"a", "one"}, data{"b", "two"}, data{"c", "three"}, data{"d", "four"}, data{"e", "five"}}

	db, err := pebble.Open("", &pebble.Options{FS: vfs.NewMem()})
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range d {
		key := []byte(v.key)
		if err := db.Set(key, []byte(v.value), pebble.Sync); err != nil {
			log.Fatal(err)
		}
	}

	value, closer, err := db.Get([]byte("a"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %s\n", "a", value)
	if err := closer.Close(); err != nil {
		log.Fatal(err)
	}
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
	// Output:
	// hello world
}
