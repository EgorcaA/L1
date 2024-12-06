package main

import (
	"context"
	"fmt"
	"time"
)

func sender(ctx context.Context, ch chan int) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			// Завершаем отправку данных при истечении времени
			fmt.Println("Sender: stopping...")
			close(ch) // Закрываем канал
			return
		case ch <- i:
			// Отправляем данные в канал
			fmt.Printf("Sender: sent %d\n", i)
			i++
			time.Sleep(500 * time.Millisecond) // Искусственная задержка
		}
	}
}

func receiver(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			// Завершаем чтение данных при истечении времени
			fmt.Println("Receiver: stopping...")
			return
		case value, ok := <-ch:
			if !ok {
				// Канал закрыт, завершаем работу
				fmt.Println("Receiver: channel closed, exiting...")
				return
			}
			// Читаем данные из канала
			fmt.Printf("Receiver: received %d\n", value)
		}
	}
}

func main() {
	// Время работы программы
	N := 5 * time.Second

	// Создаем канал
	ch := make(chan int)

	// Контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), N)
	defer cancel()

	// Запуск отправителя
	go sender(ctx, ch)

	// Запуск получателя
	go receiver(ctx, ch)

	// Ждем истечения времени
	<-ctx.Done()
	fmt.Println("Main: program finished")
}
