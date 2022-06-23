package main

import "testing"

func TestSome(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "test1", want: "some"},
		{name: "test2", want: "thing"},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			t.Logf("%v\n", v.want)
		})
	}

}
