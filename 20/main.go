// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// Разбиваем строку на слова
	words := strings.Fields(s)

	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Собираем слова обратно в строку
	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)

	fmt.Println("Оригинал:", input)
	fmt.Println("Перевернутый:", reversed)
}
