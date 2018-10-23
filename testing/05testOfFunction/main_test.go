package main

import (
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestGetSome(t *testing.T) {
	t.Log("Testing the getSome function in main")

	if !getSome(false) {
		t.Errorf("%v : getSome failed ", failed)
	} else {
		t.Logf("%v : getSome succeded ", succeed)
	}
}
