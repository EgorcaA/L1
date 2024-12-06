package main

import (
	"fmt"
	"sync"
)

func calculateSquare(number int, wg *sync.WaitGroup, results chan int) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении горутины
	square := number * number
	results <- square // Отправляем результат в канал
}

func main() {
	// Массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Канал для передачи результатов
	results := make(chan int, len(numbers))

	// WaitGroup для синхронизации горутин
	var wg sync.WaitGroup

	// Запускаем горутины для вычисления квадратов
	for _, num := range numbers {
		wg.Add(1) // Увеличиваем счетчик горутин
		go calculateSquare(num, &wg, results)
	}

	// Ждем завершения всех горутин
	wg.Wait()
	close(results) // Закрываем канал после завершения всех вычислений

	// Выводим результаты
	for square := range results {
		fmt.Println(square)
	}
}
