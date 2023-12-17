package main

import (
	"cmp"
	"fmt"
	"strings"
)

// УСЛОВИЕ:
// Реализовать бинарный поиск встроенными методами языка.

// LeftBinSearchImpl левый бинарный поиск, первое подходящее значение (самый левый элемент data для которого checkFunc дает true)
func LeftBinSearchImpl[T cmp.Ordered](checkFunc func(value T) bool, data []T) T {
	var m int
	l, r := 0, len(data)-1
	for l < r {
		m = (l + r) / 2
		if checkFunc(data[m]) {
			r = m
		} else {
			l = m + 1
		}
	}
	return data[l]
}

// RightBinSearchImpl последнее подходящее значение (самый правый элемент data для которого checkFunc дает true)
func RightBinSearchImpl[T cmp.Ordered](checkFunc func(value T) bool, data []T) T {
	var m int
	l, r := 0, len(data)-1
	for l < r {
		m = (l + r + 1) / 2
		if checkFunc(data[m]) {
			l = m
		} else {
			r = m - 1
		}
	}
	return data[l]
}

func main() {
	// Бинарный поиск предполагает поиск по отсортированному набору данных

	// case 1: найти самый левый элемент, который >= 5 и < 8 (Левый бинарный поиск)
	case1data := []int{1, 2, 4, 5, 6, 7, 9, 10, 30, 90}
	fmt.Println(LeftBinSearchImpl(func(val int) bool {
		if val >= 5 && val < 8 {
			return true
		}
		return false
	},
		case1data))

	// case 2: найти самый правый элемент, который >= 5 и < 8 (Правый бинарный поиск)
	case2data := []int{1, 2, 4, 5, 6, 7, 9, 10, 30, 90}
	fmt.Println(RightBinSearchImpl(func(val int) bool {
		if val >= 5 && val < 8 {
			return true
		}
		return false
	},
		case2data))

	// case 3: Поиск человека близкого к имени "Henry" в телефонном справочнике (Henry или следующего имени за ним)
	// для простоты только ASCII символы
	case3data := []string{"Alice", "Benjamin", "Chloe", "David", "Emily", "Frederick", "Grace", "Henry", "Isabella", "Jack", "Katherine", "Liam", "Madison", "Nathan", "Olivia", "Peter", "Quinn", "Rachel", "Samuel", "Taylor", "Ulysses", "Victoria", "William", "Xavier", "Yasmine", "Zachary"}
	fmt.Println(LeftBinSearchImpl(func(val string) bool {
		if strings.Compare(val, "Henry") >= 0 { // лексикографически сравнивает идет ли строка val дальше по алфавиту от Henry (включительно)
			return true
		}
		return false
	},
		case3data))
}
