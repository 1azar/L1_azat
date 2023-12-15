package main

import "fmt"

// УСЛОВИЕ:
//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Human структура с одним полем - Name, и методоми с рессивером по значению - Greet, GetStructName
type Human struct {
	Name string
}

func (h Human) Greet() {
	fmt.Println("Hi! I'm " + h.Name)
}

func (h Human) GetStructName() {
	fmt.Println("Human")
}

// Action структура, "наследующая" за счет композиции, поля и методы Human.
// С собственным методом GetStructName, которая переписывает одноименный родительский метод
type Action struct {
	Human // встраивание структуры Human в дочернюю структуру Action
}

func (a Action) GetStructName() {
	fmt.Println("Action")
}

func main() {

	// создание экземпляра Human, логирование экспортированного поля Name и вызов его методов
	var h1 Human = Human{Name: "Alice"}
	fmt.Println(h1.Name) // "Alice"
	h1.Greet()           // "Hi! I'm Alice"
	h1.GetStructName()   // "Human"

	fmt.Println("") // разделитель

	// создание экземпляра Action, логирование поля Name и вызов его методов
	a1 := Action{} // При такой инициализации нельзя определить поле Name внутри {}. Name = ""
	a1.Name = "Meta Human"
	fmt.Println(a1.Name) // обращение к полю "наследованной" структурой Action от предка Human. "Meta Human"
	a1.Greet()           // вызов "наследованного" метода. "Hi! I'm Meta Human"
	a1.GetStructName()   // данный метод имеют, и потомок Action, и родитель Human. Метод потомка перепишет родительский. "Action"

	fmt.Println("") // разделитель

	// Создание экземпляра Action на основе объекта h1. Значения поля будут скопированы с h1.
	a2 := Action{h1}
	fmt.Println(a2.Name) // "Alice"
	a2.Greet()           // "Hi! I'm Alice"
	a2.GetStructName()   // "Action"
}
