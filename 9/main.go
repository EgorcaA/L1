package main

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

import (
	"fmt"
)

func sendNumbers(ch chan int, numbers []int) {
	// for i := 0; i < count; i++ {
	for _, num := range numbers {
		ch <- num
	}
	close(ch)
}

func multiply(chInp chan int, chOut chan int) {
	var num int
	for num = range chInp {
		chOut <- num * 2
	}
	close(chOut)
}

func main() {
	// fmt.Printf("hello\n")
	numbers := []int{1, 2, 3}
	var ch1 chan int = make(chan int)
	var ch2 chan int = make(chan int)

	go sendNumbers(ch1, numbers)
	go multiply(ch1, ch2)

	for result := range ch2 {
		fmt.Println(result)
	}
}
