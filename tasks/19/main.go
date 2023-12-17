package main

import "fmt"

// УСЛОВИЕ:
// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.

// ReverseString переворачивает строку с учетом Unicode кодировки
func ReverseString(input string) string {
	// Преобразовать строку в срез рун
	runes := []rune(input)

	// Получить длину строки
	length := len(runes)

	// Инвертировать срез рун
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Создать строку из инвертированного среза рун
	return string(runes)
}

func main() {
	fmt.Print("Введите строку: ")

	// Считать строку с консоли
	var input string
	fmt.Scanln(&input)

	// Перевернуть строку с учетом Unicode
	result := ReverseString(input)

	// Вывести результат
	fmt.Printf("Перевернутая строка: %s\n", result)
}
