// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

// Структура Point с инкапсулированными полями
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Геттеры для доступа к полям x и y
func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

// Функция для вычисления расстояния между двумя точками
func (p1 *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.GetX()-p1.GetX(), 2) + math.Pow(p2.GetY()-p1.GetY(), 2))
}

func main() {
	// Создаем две точки с помощью конструктора
	point1 := NewPoint(3.0, 4.0)
	point2 := NewPoint(7.0, 1.0)

	fmt.Printf("Точка 1: (%v, %v)\n", point1.GetX(), point1.GetY())
	fmt.Printf("Точка 2: (%v, %v)\n", point2.GetX(), point2.GetY())

	// Вычисляем и выводим расстояние между точками
	dist := point1.Distance(point2)
	fmt.Printf("Расстояние между точками: %.2f\n", dist)
}
