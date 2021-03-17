package main

import (
	"strings"
)

type flagStringSlice struct {
	value  string
	ok     bool
	values []string
}

func (f *flagStringSlice) String() string {
	return ""
}

func (f *flagStringSlice) Set(s string) error {
	f.value = s
	f.Parse()
	return nil
}

func (f *flagStringSlice) Parse() error {
	if len(f.value) == 0 {
		return nil
	}

	fv := f.value
	sp := strings.Split(fv, ",")
	f.ok = true
	f.values = sp
	return nil
}
