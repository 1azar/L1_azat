package main

import "fmt"

// УСЛОВИЕ:
// Поменять местами два числа без создания временной переменной.

func main() {
	a := 420
	b := 419
	fmt.Printf("a=%d\tb=%d\n", a, b)
	a, b = b, a
	fmt.Printf("a=%d\tb=%d\n", a, b)
}
