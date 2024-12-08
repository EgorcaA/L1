package main

import (
	"fmt"
)

// Реализовать пересечение двух неупорядоченных множеств.

func checkIntersect(set1, set2 []int) (intersection []int) {

	//хотим искать ключи в словареб структура просто места не занимает
	elements1 := make(map[int]struct{})
	for _, elem := range set1 {
		elements1[elem] = struct{}{}
	}
	for _, elem2 := range set2 {
		if _, exists := elements1[elem2]; exists {
			intersection = append(intersection, elem2)
		}
	}

	return
}

func main() {
	set1 := []int{1, 2, 3}
	set2 := []int{2, 3, 4}
	intersection := checkIntersect(set1, set2)
	fmt.Println(intersection)
}
