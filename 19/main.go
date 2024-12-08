package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

func reverseString(s string) string {
	// Преобразуем строку в срез рукописных символов (rune), чтобы корректно работать с Unicode символами
	runes := []rune(s)

	// Инвертируем срез
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	// Пример строки, содержащей Unicode символы
	input := "главрыба"
	reversed := reverseString(input)

	fmt.Println("Оригинал:", input)
	fmt.Println("Перевернутый:", reversed)
}
