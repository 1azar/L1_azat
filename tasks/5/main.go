package main

import (
	"context"
	"fmt"
	"time"
)

// УСЛОВИЕ:
// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	// Значение таймаута в секундах. Длительность работы программы
	N := 10 * time.Second

	// корневой контекст
	ctx := context.Background()

	// создание контекста с таймаутом.
	ctx, cancel := context.WithTimeout(ctx, N)
	defer cancel() // Отложено отменит контекст при завершении main. Не понадобится, но пусть будет
	startT := time.Now()

	// Канал в который будут записываться данные.
	dataChan := make(chan any, 10)

	// Горутина отправляющая значения в канал
	go func(ctx context.Context, ch chan<- any) {
		for i := 1; ; i++ {
			select {
			case <-ctx.Done(): // Тут канал в который придет сигнал когда контекст отменится (таймаут выйдет или cancel())
				close(ch) // закрытие канала
				return    // Прерывание выполнения если контекст закроется
			case <-time.Tick(500 * time.Millisecond): // отправка в канал сообщения через dt
				ch <- fmt.Sprintf("message\t%d", i)
			}
		}
	}(ctx, dataChan)

	// Чтение записей из канала
	for msg := range dataChan {
		fmt.Printf("received: %s\n", msg)
	}

	fmt.Println("Главный поток завершился. Период работы ~ ", time.Since(startT))

}
