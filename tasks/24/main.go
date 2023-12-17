package main

import (
	"fmt"
	"math"
	"strings"
)

// УСЛОВИЕ:
// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

// Point - структура для представления точки в двумерном пространстве
type Point struct {
	x, y float64
}

// NewPoint - конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Distance - функция для вычисления расстояния между двумя точками
func Distance(p1, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1, 2)
	point2 := NewPoint(40, 6)

	// Вычисляем расстояние между точками
	distance := Distance(point1, point2)

	// Вывод результата
	fmt.Printf("Расстояние между точкой (%.2f, %.2f) и точкой (%.2f, %.2f) равно %.2f\n",
		point1.x, point1.y, point2.x, point2.y, distance)
	fmt.Printf("p1%s%.2f%sp2", strings.Repeat("-", int(distance/2)), distance, strings.Repeat("-", int(distance/2)))
}
