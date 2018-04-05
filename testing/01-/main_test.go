//Testing package for main. This file only contains the functions,
//and no main function
package main

import (
	"testing"
)

/*
func TestSum(t *testing.T) {
	if sum(5, 5) != 10 {
		t.Errorf("The sum was not correct \n")
	} else {
		t.Logf("The test exexuted correctly \n")
	}
}
*/

func TestGetWeb1(t *testing.T) {
	resp, err := getWeb("https://erter.org")

	if resp.Status == "200 OK" {
		t.Logf("Loading web worked OK, %v \n", err)
	} else {
		t.Errorf("Test failed loading web : %v \n", err)
	}
}

func TestGetWeb2(t *testing.T) {
	resp, err := getWeb("https://erter.org/thisPageDoesNotExist.html")

	if resp.Status != "200 OK" {
		//t.Fail()
		t.Errorf("Loading the page failed with err = %v, %v \n", err, resp.Status)
	} else {
		t.Logf("The page loaded perfectly \n")
	}
}
