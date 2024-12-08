// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

package main

import (
	"fmt"
	"strings"
)

func areAllUnique(str string) bool {
	// Приводим строку к нижнему регистру
	str = strings.ToLower(str)

	// Создаем мапу для отслеживания встреченных символов
	seen := make(map[rune]bool)

	// Проходим по каждому символу строки
	for _, char := range str {
		if seen[char] { // Если символ уже был, возвращаем false
			return false
		}
	}

	// Если все символы уникальны
	return true
}

func main() {
	str1 := "abcde"
	str2 := "aAbCdE"
	str3 := "hello"
	str4 := "world!"

	fmt.Println("Все символы уникальны (str1):", areAllUnique(str1)) // true
	fmt.Println("Все символы уникальны (str2):", areAllUnique(str2)) // true
	fmt.Println("Все символы уникальны (str3):", areAllUnique(str3)) // false
	fmt.Println("Все символы уникальны (str4):", areAllUnique(str4)) // true
}
