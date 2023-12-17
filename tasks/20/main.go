package main

import (
	"fmt"
	"strings"
)

// УСЛОВИЕ:
// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

// ReverseWords переворачивает слова в строке
func ReverseWords(input string) string {
	// Fields разбивает строку input вокруг каждого экземпляра одного или нескольких последовательных символов пробелов,
	//как это определено в unicode.IsSpace, возвращающий фрагмент подстрок input или пустой фрагмент, если s содержит только пробелы.
	words := strings.Fields(input)

	// Получить длину строки
	wordsCount := len(words)

	// Инвертировать срез строк
	for i, j := 0, wordsCount-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(ReverseWords("snow dog sun"))
}
