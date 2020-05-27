package main

import (
	"fmt"
	"time"
)

func run() {
	// Modify the code so you execute task1 and task2
	// so when you run go test -v you should see both messages
}

func task1() {
	time.Sleep(2 * time.Second)
	fmt.Println("Finished task 1")
}

func task2() {
	time.Sleep(3 * time.Second)
	fmt.Println("Finished task 2")
}