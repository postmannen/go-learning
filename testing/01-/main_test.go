//Testing package for main. This file only contains the functions,
//and no main function
package main

import (
	"testing"
)

func TestAddByTwo(t *testing.T) {
	if addByTwo(4) != 6 {
		t.Error("addByTwo gave the wrong result")
	} else {
		t.Log("adding by two ok")
	}
}
