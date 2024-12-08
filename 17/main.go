package main

// Реализовать бинарный поиск встроенными методами языка.

import "fmt"

// Функция бинарного поиска
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	// Пока левая граница не пересекает правую
	for left <= right {
		// Находим средний индекс
		mid := left + (right-left)/2

		// Проверяем элемент на среднем индексе
		if arr[mid] == target {
			return mid // Элемент найден, возвращаем индекс
		} else if arr[mid] < target {
			// Если элемент в середине меньше искомого, ищем в правой части
			left = mid + 1
		} else {
			// Если элемент в середине больше искомого, ищем в левой части
			right = mid - 1
		}
	}

	// Если элемент не найден, возвращаем -1
	return -1
}

func main() {
	// Пример массива для сортировки
	arr := []int{10, 7, 8, 9, 1, 5}
	target := 8
	fmt.Println("массив:", arr)

	index := binarySearch(arr, target)

	// Выводим результат
	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d\n", target, index)
	} else {
		fmt.Println("Элемент не найден")
	}
}
