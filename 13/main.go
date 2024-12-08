// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Printf("Before: a = %d, b = %d\n", a, b)

	// Меняем местами
	a = a ^ b
	//0101 a
	//0011 b
	//0110 a ^ b

	b = a ^ b
	//0110 a
	//0011 b
	//0101 a ^ b

	a = a ^ b
	//0110 a
	//0101 b
	//0011 a ^ b

	fmt.Printf("After: a = %d, b = %d\n", a, b)
}
