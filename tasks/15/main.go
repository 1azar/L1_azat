package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

// УСЛОВИЕ:
//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
//
//
//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	/*
		Не стоит работать со строкой как с массивом тк есть риск обрезать строку посередине многобайтового символа.

		Строка - неизменяемый массив байт, при этом каждый символ будет закодирован в байте только для ASCII (abcdef..12345..).
		Соответственно, если createHugeString() вернет символы из кириллицы или любые другие символы, для кодировки которых требуется больше места чем один байт (明 珠 ...)
		Взятие среза таким образом [:100] приведет к нарушению кодировки и мы потеряем символы. Например: Вместо "абв" получим "а�" если брать "абвгде"[:3].
		В Go строки представляются в UTF-8 (для одного символа выделяется от 1го до 4х байт), поэтому для получения подстроки нужно учитывать сколько байт взять из массива для
		дальнейшего декодирования по таблице Unicode.
		Для работы со строкой необходимо использовать тип rune(int32) и пакет unicode/utf8.
	*/
	justString2 := string([]rune(v)[:100])

	fmt.Println("(НЕВЕНРНО) justString = v[:100]: ", justString)
	fmt.Println("(ПРАВИЛЬНО) justString2 := string([]rune(v)[:100]): ", justString2)

	// выведет символы и сколько байт они занимают
	str := "Hello, 世界, Мир! "
	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)
		fmt.Printf("%c %v\n", r, size)
		str = str[:len(str)-size]
	}

}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	//return string(bytes.Repeat([]byte("x"), size))
	return string(bytes.Repeat([]byte("識"), size))
}
