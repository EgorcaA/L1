package main

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

import "fmt"

// Функция для быстрой сортировки
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		// Если массив содержит 0 или 1 элемент, он уже отсортирован
		return arr
	}

	// Выбираем опорный элемент (в данном случае первый элемент)
	pivot := arr[0]

	// Разделяем массив на два среза: один с элементами меньшими pivot, другой - с элементами большими pivot
	var left, right []int
	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	// Рекурсивно сортируем левую и правую части массива и соединяем их с опорным элементом
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	// Пример массива для сортировки
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Исходный массив:", arr)

	// Сортируем массив
	sortedArr := quickSort(arr)

	// Выводим отсортированный массив
	fmt.Println("Отсортированный массив:", sortedArr)
}
