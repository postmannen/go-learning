package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzConvert(f *testing.F) {
	f.Add("ABCD")

	f.Fuzz(func(t *testing.T, orig string) {
		r := Reverse(orig)
		rr := Reverse(r)

		//t.Log(s)

		if utf8.ValidString(orig) && !utf8.ValidString(r) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", r)
		}

		if orig != rr {
			t.Errorf("double reverse failed: Orig: %x, double reverse: %x\n", orig, rr)
		}
	})
}
