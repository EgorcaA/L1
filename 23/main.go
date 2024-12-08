// Удалить i-ый элемент из слайса.

package main

import "fmt"

func removeElement(slice []int, index int) []int {

	if index < 0 || index >= len(slice) {
		return slice // Возвращаем оригинальный слайс, если индекс некорректен
	}

	// Срезаем слайс, исключая элемент по индексу
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("Исходный слайс:", slice)
	indexToRemove := 2
	slice = removeElement(slice, indexToRemove)

	fmt.Println("После удаления элемента с индексом", indexToRemove, ":", slice)
}
