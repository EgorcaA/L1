package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Worker function to process data
func worker(id int, dataChan <-chan int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			// Exit on signal
			fmt.Printf("Worker %d exiting...\n", id)
			return
		case data, ok := <-dataChan:
			if !ok {
				// Channel closed, exit
				fmt.Printf("Worker %d: data channel closed, exiting...\n", id)
				return
			}
			// Process data
			fmt.Printf("Worker %d processing data: %d\n", id, data)
		}
	}
}

func main() {
	// Number of workers (could be passed as a flag or input)
	numWorkers := 4

	// Channels
	dataChan := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup

	// Context for handling Ctrl+C
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, done, &wg)
	}

	// Main loop: continuously send data into the channel
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(done)     // Signal workers to stop
				close(dataChan) // Close data channel
				return
			default:
				dataChan <- rand.Intn(100) // Send random data
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Wait for termination signal
	<-ctx.Done()
	fmt.Println("Shutting down gracefully...")
	wg.Wait() // Wait for all workers to finish
	fmt.Println("All workers exited. Program terminated.")
}
