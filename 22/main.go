// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a и b, значение которых > 2^20.

package main

import (
	"fmt"
)

// за int64 вылезти сложно
func main() {
	a := int64(1048577) // Число больше 2^20
	b := int64(2097152)

	fmt.Println("Число a:", a)
	fmt.Println("Число b:", b)

	// Сложение
	sum := a + b
	fmt.Println("Сложение: a + b =", sum)

	// Вычитание
	diff := a - b
	fmt.Println("Вычитание: a - b =", diff)

	// Умножение
	product := a * b
	fmt.Println("Умножение: a * b =", product)

	// Деление
	if b != 0 {
		quotient := a / b
		fmt.Println("Деление: a / b =", quotient)
	} else {
		fmt.Println("Ошибка: деление на ноль")
	}
}
