package main

import (
	"fmt"
	"math"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func groupTemps(temps []float64) map[int][]float64 {
	groupedTemps := make(map[int][]float64)
	var keyT int
	for _, T := range temps {
		keyT = int(math.Floor(T/10.0)) * 10
		groupedTemps[keyT] = append(groupedTemps[keyT], T)
	}
	return groupedTemps
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groupedTemps := groupTemps(temps)

	for keyT, vals := range groupedTemps {
		fmt.Printf("%d: %v\n", keyT, vals)
	}
}
