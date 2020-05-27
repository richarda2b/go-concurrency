package main

import (
	"testing"
	"time"
)

func TestRunningTime(t *testing.T) {

	start := time.Now()
	run()

	if time.Since(start) > (4 * time.Second) {
		t.Fatal("Run() took too long")
	}
}
