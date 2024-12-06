package main

import (
	"fmt"
	"time"
)

func worker() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Worker stopped by panic:", r)
		}
	}()
	for {
		fmt.Println("Worker is running...")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go worker()

	time.Sleep(2 * time.Second)
	panic("Stop worker")
	time.Sleep(1 * time.Second)
}
