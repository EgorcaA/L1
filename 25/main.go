// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	// канал закроется и рутина разблокируется через время
	<-time.After(duration)
}

func main() {
	fmt.Println("Засыпаем на 2 секунды...")
	mySleep(2 * time.Second)
	fmt.Println("Проснулись")
}
