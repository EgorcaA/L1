package main

import (
	"fmt"
)

func setBit(num int64, i uint, value bool) int64 {
	if value {
		// Установить i-й бит в 1
		return num | (1 << i) // побитово или 2 числа int64, справа сдвиг влево
	}
	// Установить i-й бит в 0
	return num & ^(1 << i)
}

func main() {
	var num int64 = 42 // Пример числа: 42 (в двоичной форме 101010)
	fmt.Printf("Initial value: %d (binary: %b)\n", num, num)

	// Установить 3-й бит в 1
	num = setBit(num, 3, true)
	fmt.Printf("After setting 3rd bit to 1: %d (binary: %b)\n", num, num)

	// Установить 1-й бит в 0
	num = setBit(num, 1, false)
	fmt.Printf("After setting 1st bit to 0: %d (binary: %b)\n", num, num)
}
