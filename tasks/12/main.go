package main

import "fmt"

// УСЛОВИЕ:
//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {

	data := []string{"cat", "cat", "dog", "cat", "tree"}

	// 1 способ с использованием мапы.
	set1 := []string{}
	hash := make(map[string]struct{})

	for _, v := range data {
		hash[v] = struct{}{} // Если ключа нет - он добавится, если есть обновится. Дублирующихся ключей не бывает в мапе
	}
	// Ключи hash - множество уникальных значений
	for key := range hash {
		set1 = append(set1, key)
	}

	fmt.Println(set1)

	// 2 способ с проверкой в слайсе
	set2 := []string{}
	// Функция для проверки есть ли значению в слайсе. Передача по значению тк копирования данных все равно не произойдет как с массивом
	isInSlice := func(slice []string, val string) bool {
		for _, v := range slice {
			if v == val {
				return true
			}
		}
		return false
	}
	for _, v := range data {
		if isInSlice(set2, v) { // если элемент уже есть в множестве: пропускаем без добавления
			continue
		}
		set2 = append(set2, v)
	}

	fmt.Println(set2)

}
