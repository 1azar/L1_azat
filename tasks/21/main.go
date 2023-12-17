package main

import (
	"fmt"
	"math/rand"
)

// УСЛОВИЕ:
// Реализовать паттерн «адаптер» на любом примере.

// NameGenerator интерфейс. Если аргументу функции, или полю структуры, или переменной назначается этот тип при объявлении,
//
//	то инициализировать их можно любым объектом имеющим перечисленные в интерфейсе методы
type NameGenerator interface {
	GenerateName() string
}

// RusNames реализует NameGenerator интерфейс
type RusNames struct{}

func (rn RusNames) GenerateName() string {
	names := []string{"Ваня", "Алиса", "Саша", "Женя"}
	return names[rand.Intn(len(names)-1)]
}

// EngNames реализует NameGenerator интерфейс
type EngNames struct{}

func (en EngNames) GenerateName() string {
	names := []string{"Bob", "Alice", "Steve", "Harry"}
	return names[rand.Intn(len(names)-1)]
}

// CharacterDesigner имеет поле, где тип определен интерфейсом, значит этому полю можно присвоить объекты реализующие интерфейс.
// Это позволяет не изменять существующий код, просто писать новую реализацию и подключить ее при инициализации структуры (вставить в адаптер)
type CharacterDesigner struct {
	GetRandomName NameGenerator // Этом полю можно присвоить любой объект имеющий метод GenerateName() string
}

func main() {

	// При создании объекта CharacterDesigner в поле, где тип определен интерфейсом можно присвоить любой объект реализующий данный интерфейс
	myEngDesigner := CharacterDesigner{GetRandomName: RusNames{}}
	myRusDesigner := CharacterDesigner{GetRandomName: EngNames{}}

	fmt.Println(myEngDesigner.GetRandomName.GenerateName())
	fmt.Println(myEngDesigner.GetRandomName.GenerateName())
	fmt.Println(myEngDesigner.GetRandomName.GenerateName())
	fmt.Println(myRusDesigner.GetRandomName.GenerateName())
	fmt.Println(myRusDesigner.GetRandomName.GenerateName())
	fmt.Println(myRusDesigner.GetRandomName.GenerateName())

}
