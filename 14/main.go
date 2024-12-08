package main

import "fmt"

//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

func determineType(v interface{}) {
	switch v.(type) {
	// работает только для интерфейсов в свитче
	// вообще предразование типа int(var)
	case int:
		fmt.Println("Type is int")
	case string:
		fmt.Println("Type is string")
	case bool:
		fmt.Println("Type is bool")
	case chan int:
		fmt.Println("Type is channel of int")
	default:
		fmt.Println("Unknown type")
	}

	// if value, ok := v.(int); ok {
	// 	fmt.Printf("Value is int: %d\n", value)
	// }
}

func main() {
	// Примеры переменных разных типов
	var a int = 42
	var b string = "hello"
	var c bool = true
	var d chan int = make(chan int)

	// Проверяем типы переменных
	determineType(a)
	determineType(b)
	determineType(c)
	determineType(d)
	determineType(3.14) // Неизвестный тип
}
