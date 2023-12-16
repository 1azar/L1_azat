package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// УСЛОВИЕ:
// Реализовать все возможные способы остановки выполнения горутины.

func main() {

	// СПОСОБ 1: Есть канал в который по необходимости что-то кладется, горутина в цикле проверяет канал и если там что-то есть -> завершает работу.
	// Сигнал будет для одной горутины - той, которая первым заберет запись из канала. Тк в этой задаче только 1 горутина это не важно.
	// Чтобы дождаться горутину используем wg
	var wg1 sync.WaitGroup
	// Канал куда отправится сигнал для завершения. Размер struct{} = 0 байт, поэтому такой тип в канале - эффективно
	ch1 := make(chan struct{})
	wg1.Add(1)                  // +1 счетчику WaitGroup
	go Foo1(ch1, &wg1)          // старт горутины
	time.Sleep(4 * time.Second) // спим 4 сек и отправляем сигнал для завершения горутины
	ch1 <- struct{}{}
	wg1.Wait() // Ждем пока горутина завершится.

	// СПОСОБ 2: Есть канал, который когда надо закроют - это будет сигналом завершения работы.
	// Тут сигнал распространится на все горутины, они все завершат работу.
	// Чтобы дождаться горутину используем wg
	var wg2 sync.WaitGroup
	// Канал, который закроется при необходимости завершения горутины
	ch2 := make(chan struct{})
	wg2.Add(1)                  // +1 счетчику WaitGroup
	go Foo1(ch2, &wg2)          // старт горутины
	time.Sleep(4 * time.Second) // спим 4 сек и закрываем канал
	close(ch2)
	wg2.Wait() // Ждем пока горутина завершится

	// СПОСОБ 3: Отправляем в горутину указатель на общую переменную, когда нужно остановить работу - меняем значение этой переменной.
	// Чтобы дождаться горутину используем wg
	var wg3 sync.WaitGroup
	// Общая переменная
	stopGoroutine := false
	wg3.Add(1)                    // +1 счетчику WaitGroup
	go Foo2(&stopGoroutine, &wg3) // старт горутины
	time.Sleep(4 * time.Second)   // спим 4 сек и меняем значение переменной указатель на которую есть у горутины
	stopGoroutine = true
	wg3.Wait() // Ждем пока горутина завершится

	// СПОСОБ 4: Используем что-то, что реализует интерфейс context.Context - структура, которая дает тот же канал, но через метод Done(). Плюс дополнительные функции.
	// Чтобы дождаться горутину используем wg
	var wg4 sync.WaitGroup
	// Получаем сам контекст и функцию которая отменяет его (если нам надо отменить его до конца таймаута, мало ли)
	// Создается на основе родительского контекста, если предок отменен, то и он отменится. Отмена распространится по всему дереву контекстов
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	wg4.Add(1)         // +1 счетчику WaitGroup
	go Foo4(ctx, &wg4) // ctx принято всегда первым аргументом в функцию отправлять
	time.Sleep(2 * time.Second)
	cancel()   // отменяем контекст раньше истечения таймаута
	wg4.Wait() // Ждем пока горутина завершится

	// Способ 5: Использование объекта sync.Cond. У нее есть методы для закрытия одной или всех горутин:
	// Из документации: Broadcast corresponds to closing a channel, and Signal corresponds to sending on a channel
	var (
		mu   sync.Mutex
		cond *sync.Cond = sync.NewCond(&mu)
	)
	// Запускаем горутину
	go func(mu *sync.Mutex, cond *sync.Cond) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println("Foo5: started")
		go func() {
			for {
				fmt.Println("Foo5: doing cool things that takes 999 hours")
				time.Sleep(999 * time.Hour)
			}
		}() // не передаю аргументы тк функция анонимная может захватить переменные внешней функции main
		// Ждем сигнала
		// Преимущество данного способа, перед предыдущими в том, что если полезная работа очень долгая,
		// то при сигнале закрытия выполнение прервется, а не "не начнется"
		cond.Wait()
		fmt.Println("Foo5: stopped")
	}(&mu, cond)

	time.Sleep(4 * time.Second) // Спим 4 секунды
	mu.Lock()
	cond.Signal() // Сигнал завершения горутины
	mu.Unlock()
	time.Sleep(time.Second) // подождем пока горутина полностью выполнит логику после cond.Wait() иначе main завершится раньше

	// Все способы отмены выполнения горутин завязаны на разделении одного общего ресурса (переменная, канал, контекст).
	// Контекст это суперсет над каналом, а канал суперсет над указателем на общую переменную

}

// Foo1 использует канал для сигнализирования о завершении
func Foo1(stopFun <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Foo1: started")
	for {
		select {
		default: // тут полезная работа
			fmt.Println("Foo1: doing cool things")
			time.Sleep(time.Second)
		case <-stopFun: // Конец веселью. Сворачиваемся. В канал либо что-то записали, либо его закрыли
			fmt.Println("Foo1: stopped")
			return
		}
	}
}

// Foo2 использует указатель на переменную для сигнализирования о завершении
func Foo2(stopFun *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Foo2: started")
	for {
		if *stopFun { // если true: закончить все это
			fmt.Println("Foo2: stopped")
			return
		}
		// работаем, работаем, работаем
		fmt.Println("Foo2: doing cool things")
		time.Sleep(time.Second)
	}
}

// Foo4 использует контекст (который внутри содержит все тот же канал)
func Foo4(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Foo4: started")
	for {
		select {
		default: // тут полезная работа
			fmt.Println("Foo4: doing cool things")
			time.Sleep(time.Second)
		case <-ctx.Done(): // Конец веселью. Сворачиваемся.
			fmt.Println("Foo4: stopped")
			return
		}
	}
}
