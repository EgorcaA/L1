package main

import (
	"fmt"
	"sync"
	"time"
)

// Структура, оборачивающая карту и мьютекс
type SafeMap struct {
	mu   sync.Mutex
	data map[int]int
}

func (sm *SafeMap) Set(key, value int) {
	sm.mu.Lock()         // Блокируем доступ к map
	defer sm.mu.Unlock() // Освобождаем блокировку после записи
	sm.data[key] = value
}

func (sm *SafeMap) Get(key int) int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.data[key]
}

func worker(id int, sm *SafeMap, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		sm.Set(i, id*100+i)
		fmt.Printf("Worker %d: Set key %d\n", id, i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	sm := &SafeMap{data: make(map[int]int)}

	var wg sync.WaitGroup
	// Запуск нескольких горутин, которые будут записывать в карту
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, sm, &wg)
	}

	wg.Wait() // Ожидаем завершения всех горутин

	// Выводим результаты
	for i := 0; i < 5; i++ {
		fmt.Printf("Main: Key %d = %d\n", i, sm.Get(i))
	}
}
