package main

import (
	"time"
)

func run2() int {
	// Modify the code so it returns the result of A divided  by B
	// r: 5
	c := make(chan int)

	go func() {
		c <- getA()
	}()

	go func() {
		c <- getB()
	}()

	a, b := <-c, <-c
	return a / b
}

func getA() int {
	time.Sleep(2 * time.Second)
	return 25
}

func getB() int {
	time.Sleep(3 * time.Second)
	return 5
}
