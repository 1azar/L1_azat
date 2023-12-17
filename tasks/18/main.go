package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// УСЛОВИЕ:
// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

// Counter - структура счетчика
type Counter struct {
	value int64
}

// Increment - метод для инкрементации счетчика
func (c *Counter) Increment() {
	// sync/atomic обеспечивает базовые операции, такие как инкрементация и чтение, для простых типов данных
	// Можно использовать еще Mutex, RWMutex
	atomic.AddInt64(&c.value, 1)
}

// GetValue - метод для получения текущего значения счетчика
func (c *Counter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	gNums := 5
	incrCount := 100000

	fmt.Printf("%d горутин увелиичат счетчик %d раз каждый\n", gNums, incrCount)
	// Запуск нескольких горутин для инкрементации счетчика
	for i := 0; i < gNums; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrCount; j++ {
				counter.Increment()
			}
		}()
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод итогового значения счетчика
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.GetValue())
}
