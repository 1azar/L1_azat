package main

import (
	"cmp"
	"fmt"
)

// УСЛОВИЕ:
// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// partition вспомогательная функция, которая выбирает некоторый элемент массива и переставляет элементы участка массива таким образом,
// чтобы массив разбился на 2 части: левая часть содержит элементы, которые меньше этого элемента,
// а правая часть содержит элементы, которые больше или равны этого элемента.
// cmp.Ordered говорит нам что можно использовать только типы для которых поддерживают <, >, <=, >= операторы
func partition[T cmp.Ordered](a []T, l, r int) int {
	pivot := a[r]
	less := l
	for i := l; i < r; i++ {
		if a[i] <= pivot {
			a[i], a[less] = a[less], a[i]
			less++
		}
	}
	a[r], a[less] = a[less], a[r]
	return less
}

// MyQuickSort Сначала из массива нужно выбрать один элемент — его обычно называют опорным.
// Затем другие элементы в массиве перераспределяют так, чтобы элементы меньше опорного оказались до него,
// а большие или равные — после. А дальше рекурсивно применяют первые два шага к подмассивам справа и слева от опорного значения.
func quickSortImpl[T cmp.Ordered](a []T, l, r int) {
	if l < r {
		q := partition(a, l, r)
		quickSortImpl(a, l, q-1)
		quickSortImpl(a, q+1, r)
	}
}

// MyQuickSort обертка над quickSortImpl чтобы не задавать l и r при вызове
func MyQuickSort[T cmp.Ordered](a []T) {
	quickSortImpl(a, 0, len(a)-1)
}

func main() {
	a1 := []int{7, 3, 5, 4, 6, 2, 1}
	fmt.Println("'a1' before sorting: ", a1)
	MyQuickSort(a1)
	fmt.Println("'a1' sorted: ", a1)

	a2 := []string{"b", "c", "e", "a", "d"}
	fmt.Println("'a2' before sorting: ", a2)
	MyQuickSort(a2)
	fmt.Println("'a2' sorted: ", a2)

	a3 := []float64{99.0, 23, -123, -0.0001}
	fmt.Println("'a3' before sorting: ", a3)
	MyQuickSort(a3)
	fmt.Println("'a3' sorted: ", a3)

}
