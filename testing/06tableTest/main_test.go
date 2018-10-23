package main

import (
	"net/http"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2718"

func TestGet(t *testing.T) {
	tests := []struct {
		url  string
		code int
	}{
		{"https://erter.org", 200},
		{"https://erter.org/apekatt", 200},
	}

	for i, tt := range tests {
		t.Logf("\tTest %v : Testing if the http.Get methods are working.\n", i)
		resp, err := http.Get(tt.url)
		if err != nil {
			t.Errorf("\t\t%v Failed calling method http.Get %v, err = %v\n", failed, tt.url, err)
		} else {
			t.Logf("\t\t%v Succeed calling method http.Get %v , err = %v\n", succeed, tt.url, err)
		}

		t.Logf("\tTest %v : Testing if the respone code is 200.\n", i)
		if resp.StatusCode != 200 {
			t.Errorf("\t\t%v Failed status code 200 for %v, code = %v\n", failed, tt.url, resp.StatusCode)
		} else {
			t.Logf("\t\t%v Succeed status code 200 for %v, code = %v\n", succeed, tt.url, resp.StatusCode)
		}

	}

}
