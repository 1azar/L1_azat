package main

import (
	"fmt"
	"time"
)

// УСЛОВИЕ:
// Удалить i-ый элемент из слайса.

func main() {
	var st time.Time // для приблизительной оценки времени выполнения функция
	deletingIdx := 9999
	dataSize := 10000001

	// СПОСОБ 1
	a := make([]int, dataSize) //заведомо выделим нужное количество места, чтобы не было лишних аллокаций
	for i := 0; i < cap(a); i++ {
		a[i] = i
	}
	fmt.Printf("1: len=%d, cap=%d, %v, (inital state)\n", len(a), cap(a), a[deletingIdx-2:deletingIdx+3])
	st = time.Now()
	aRes1, _ := RemoveElementInSlice1(a, deletingIdx)
	fmt.Printf("1: len=%d, cap=%d, %v, (took %v)\n", len(aRes1), cap(aRes1), aRes1[deletingIdx-2:deletingIdx+3], time.Since(st))

	// СПОСОБ 2
	b := make([]int, dataSize) //заведомо выделим нужное количество места, чтобы не было лишних аллокаций
	for i := 0; i < cap(b); i++ {
		b[i] = i
	}
	fmt.Printf("2: len=%d, cap=%d, %v, (inital state)\n", len(b), cap(b), b[deletingIdx-2:deletingIdx+3])
	st = time.Now()
	aRes2, _ := RemoveElementInSlice2(b, deletingIdx)
	fmt.Printf("2: len=%d, cap=%d, %v, (took %v)\n", len(aRes2), cap(aRes2), aRes2[deletingIdx-2:deletingIdx+3], time.Since(st))

	// СПОСОБ 2
	c := make([]int, dataSize) //заведомо выделим нужное количество места, чтобы не было лишних аллокаций
	for i := 0; i < cap(c); i++ {
		c[i] = i
	}
	fmt.Printf("3: len=%d, cap=%d, %v, (inital state)\n", len(c), cap(c), c[deletingIdx-2:deletingIdx+3])
	st = time.Now()
	aRes3, _ := RemoveElementInSlice3(c, deletingIdx)
	fmt.Printf("3: len=%d, cap=%d, %v, (took %v)\n", len(aRes3), cap(aRes3), aRes3[deletingIdx-2:deletingIdx+3], time.Since(st))

}

// RemoveElementInSlice1 удаляет элемент с заданным индексом из слайса с помощью встроенной функции copy
func RemoveElementInSlice1[T any](slice []T, index int) ([]T, error) {
	// Проверка на корректность индекса
	if index < 0 || index >= len(slice) {
		return slice, fmt.Errorf("bad index value")
	}
	// Сдвигаем все элементы после удаляемого влево
	// Встроенная функция копирования копирует элементы из исходного среза в целевой срез.
	copy(slice[index:], slice[index+1:])
	// Уменьшаем длину слайса, удаляя последний элемент (дубликат)
	return slice[:len(slice)-1], nil
}

// RemoveElementInSlice2 удаляет элемент с заданным индексом из слайса с помощью встроенной функции append
func RemoveElementInSlice2[T any](slice []T, index int) ([]T, error) {
	// Проверка на корректность индекса
	if index < 0 || index >= len(slice) {
		return slice, fmt.Errorf("bad index value")
	}
	// append добавляет элементы в конец слайса (первого аргумента)
	//
	return append(slice[:index], slice[index+1:]...), nil
}

// RemoveElementInSlice3 удаляет элемент с заданным индексом из слайса, при этом МЕНЯЕТСЯ ПОРЯДОК в слайсе!!!!
// удаляемый элемент заменяется последним и возвращается первые len()-1 элементов начального слайса.
// Но если порядок не важен это наиболее эффективный способ удаления
func RemoveElementInSlice3[T any](slice []T, index int) ([]T, error) {
	// Проверка на корректность индекса
	if index < 0 || index >= len(slice) {
		return slice, fmt.Errorf("bad index value")
	}
	// элемент под удаляемым индексом теперь будет равен последнему элементу
	slice[index] = slice[len(slice)-1]
	// возвращается слайс кроме последнего элемента
	return slice[:len(slice)-1], nil
}
