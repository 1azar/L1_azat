package main

import (
	"fmt"
	"strings"
)

// УСЛОВИЕ:
// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
//abcd — true
//abCdefAaf — false
//aabcd — false

func UniqueChecking(s string) bool {
	// Приведем строку к нижнему регистру, чтобы сделать проверку регистронезависимой
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", -1) // Удалим пробелы

	// Используем map для отслеживания уникальных символов. struct{} для экономии памяти тк значения нам не нужны
	charMap := make(map[rune]struct{})

	for _, char := range []rune(s) { // для работы с символами юникод приводим все к типу рун

		if _, ok := charMap[char]; ok {
			return false // Символ уже встречался, строка не содержит уникальные символы
		}
		charMap[char] = struct{}{}
	}

	return true // Все символы уникальны
}

func main() {
	fmt.Println("ggez", " - ", UniqueChecking("ggez"))
	fmt.Println("abcdefgh", " - ", UniqueChecking("abcdefgh"))
	fmt.Println("abcdefgg", " - ", UniqueChecking("abcdefgg"))
	fmt.Println("Ало, Галя? Ты щас умрёшь!", " - ", UniqueChecking("Ало, Галя? Ты щас умрёшь!"))
	fmt.Println("Ашаган белми, тураган белә", " - ", UniqueChecking("Ашаган белми, тураган белә"))
	fmt.Println("FәЖ", " - ", UniqueChecking("FәЖ"))
	fmt.Println("Аа", " - ", UniqueChecking("Аа"))
	fmt.Println("日БВАdsaf123", " - ", UniqueChecking("日БВАdsaf123"))
	fmt.Println("x y z", " - ", UniqueChecking("x y z"))
}
