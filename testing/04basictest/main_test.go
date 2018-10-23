package main

import (
	"net/http"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestGet(t *testing.T) {
	t.Logf("Testing http.Get, and using %v and %v to show status\n", succeed, failed)

	resp, err := http.Get("https://erter.org")

	if err != nil {
		t.Errorf("%v Failed with http.Get : %v\n", failed, err)
	} else {
		t.Logf("%v Succeed with http.Get : %v\n", succeed, err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("%v Failed, wrong statuscode, should be 200 but is = %v\n", failed, resp.StatusCode)
	} else {
		t.Logf("%v Succed, statuscode is 200 = %v", succeed, resp.StatusCode)
	}

}
