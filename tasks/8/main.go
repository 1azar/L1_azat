package main

import (
	"fmt"
)

// УСЛОВИЕ
// Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.

func setBit(value int64, pos uint, bitValue bool) int64 {
	if bitValue {
		// Установка бита в 1
		// число 5 и позиция 1. 101 -> 001 << 1 = 010 -> 101 | 010 (логическое или) = 111 = 7
		// число 7 позиция 4. 00111 -> 00001<<4 = 10000 -> 00111 | 10000 = 10111 = 29
		return value | (1 << pos)
	}
	// Установка бита в 0
	return value &^ (1 << pos) // (И НЕ) если справа 1, то обнуляется левое, если 0 не меняется
}

func main() {

	var num int64
	var bit uint
	var setOne bool
	for {
		fmt.Println("Введите: <число> <позицию бита> <true/false>")
		_, err := fmt.Scan(&num, &bit, &setOne)
		if err != nil {
			fmt.Println("ошибка чтения данных: ", err)
			continue
		}
		fmt.Printf("Исходное значение %b = %d\n", num, num)
		res := setBit(num, bit, setOne)
		fmt.Printf("Результат %b = %d\n", res, res)
	}
}
