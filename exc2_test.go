package main

import (
	"testing"
	"time"
)

func TestExc2(t *testing.T) {

	start := time.Now()
	res := run2()

	if res != 5 {
		t.Fatalf("Invalid result: %v", res)
	}

	if time.Since(start) > (4 * time.Second) {
		t.Fatal("Run() took too long")
	}
}
