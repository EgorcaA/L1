package main

import (
	"fmt"
	"sync"
	"time"
)

var stop bool
var mu sync.Mutex

func worker() {
	for {
		mu.Lock()
		if stop {
			mu.Unlock()
			fmt.Println("Worker stopped by flag.")
			return
		}
		mu.Unlock()
		fmt.Println("Worker is running...")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go worker()

	time.Sleep(2 * time.Second)
	mu.Lock()
	stop = true
	mu.Unlock()

	time.Sleep(1 * time.Second)
}
