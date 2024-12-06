package main

import "fmt"

// Структура Vehicle
type Vehicle struct {
	Brand string
	Model string
}

// Метод структуры Vehicle
func (v *Vehicle) Start() {
	fmt.Printf("The %s %s is starting.\n", v.Brand, v.Model)
}

func (v *Vehicle) Stop() {
	fmt.Printf("The %s %s is stopping.\n", v.Brand, v.Model)
}

// Структура Car
type Car struct {
	Vehicle // Встраивание Vehicle
	Doors   int
}

// Метод структуры Car
func (c *Car) Honk() {
	fmt.Println("The car is honking: Beep beep!")
}

// Переопределение метода Start для Car
func (c *Car) Start() {
	fmt.Printf("The car %s %s with %d doors is starting.\n", c.Brand, c.Model, c.Doors)
}

func main() {
	// Создаем экземпляр Car
	car := Car{
		Vehicle: Vehicle{
			Brand: "Toyota",
			Model: "Corolla",
		},
		Doors: 4,
	}

	// Вызов методов
	car.Start()         // Используется переопределенный метод
	car.Honk()          // Метод Car
	car.Stop()          // Метод Vehicle через встраивание
	car.Vehicle.Start() // Вызов оригинального метода Vehicle
}
