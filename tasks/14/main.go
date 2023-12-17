package main

import (
	"fmt"
	"reflect"
)

// УСЛОВИЕ:
// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func getType(value interface{}) string {
	switch v := value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan any:
		return "channel of any" // направление канала тоже определяет его тип
	case <-chan any:
		return "reading channel of any"
	case chan<- string:
		return "writing channel of string"
	default: // если value неожидаемого типа
		// Из документации: TypeOf returns the reflection Type that represents the dynamic type of i. If i is a nil interface value, TypeOf returns nil.
		return fmt.Sprintf("unknown type (%v)", reflect.TypeOf(v))
	}
}

func main() {
	var intValue int = 42
	var stringValue string = "hello"
	var boolValue bool = true
	var inputStringChannelValue chan<- string = make(chan<- string)
	var anyChannelValue chan any = make(chan any)
	var outputAnyChannelValue <-chan any = make(<-chan any)

	fmt.Printf("Type of intValue: %s\n", getType(intValue))
	fmt.Printf("Type of stringValue: %s\n", getType(stringValue))
	fmt.Printf("Type of boolValue: %s\n", getType(boolValue))
	fmt.Printf("Type of inputStringChannelValue: %s\n", getType(inputStringChannelValue))
	fmt.Printf("Type of anyChannelValue: %s\n", getType(anyChannelValue))
	fmt.Printf("Type of outputAnyChannelValue: %s\n", getType(outputAnyChannelValue))
	fmt.Printf("For struct{foo string}: %s\n", getType(struct {
		foo string
	}{foo: "foo"}))
}
