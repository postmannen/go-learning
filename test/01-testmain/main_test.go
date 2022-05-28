package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Printf("TestMain: Start\n")

	flag.Parse()
	exitCode := m.Run()

	fmt.Printf("TestMain: End\n")
	os.Exit(exitCode)
}

func TestOne(t *testing.T) {
	fmt.Printf("TestOne\n")
}

func TestTwo(t *testing.T) {
	fmt.Printf("TestTwo\n")
}
