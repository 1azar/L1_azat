package main

import (
	"fmt"
	"sync"
)

// УСЛОВИЕ:
// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	// Создаем каналы для передачи данных между горутинами
	chanIn := make(chan int, 10)
	chanOut := make(chan int, 10)

	// Используем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Горутина для записи чисел в первый канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(chanIn) // Закрываем канал после записи всех чисел
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for _, num := range numbers {
			chanIn <- num
		}
	}()

	// Горутина для умножения чисел и записи во второй канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(chanOut)      // Закрываем канал после записи всех результатов
		for num := range chanIn { // слушает входящий канал пока все данные не прочтет
			chanOut <- num * 2
		}
	}()

	// Горутина для вывода результатов в stdout
	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range chanOut {
			fmt.Println(result)
		}
	}()
	// Ожидаем завершения всех горутин
	wg.Wait()
}
