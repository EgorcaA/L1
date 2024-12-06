package chanclose

import (
	"fmt"
)

func worker(ch chan int) {
	for val := range ch {
		fmt.Println("Received:", val)
	}
	fmt.Println("Worker stopped due to channel closure.")
}

func main() {
	ch := make(chan int)

	go worker(ch)

	ch <- 1
	ch <- 2
	close(ch) // Закрываем канал для завершения
}
