package main

import (
	"fmt"
	"time"
)

func worker(stop chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("Worker stopped.")
			return
		default:
			fmt.Println("Worker is running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stop := make(chan struct{})
	go worker(stop)

	time.Sleep(2 * time.Second)
	close(stop) // Отправляем сигнал завершения
	time.Sleep(1 * time.Second)
}
