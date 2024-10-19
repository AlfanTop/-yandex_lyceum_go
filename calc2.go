package main

import (
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Ошибка: Деление на ноль")
		return 0
	}
	return a / b
}

func main1() {
	var a, b float64
	var operator string

	fmt.Println("Введите первое число:")
	fmt.Scan(&a)

	fmt.Println("Введите оператор (+, -, *, /):")
	fmt.Scan(&operator)

	fmt.Println("Введите второе число:")
	fmt.Scan(&b)

	switch operator {
	case "+":
		fmt.Println("Результат:", add(a, b))
	case "-":
		fmt.Println("Результат:", subtract(a, b))
	case "*":
		fmt.Println("Результат:", multiply(a, b))
	case "/":
		fmt.Println("Результат:", divide(a, b))
	default:
		fmt.Println("Неверный оператор") //ввод в терминал снизу go run calc2.go
	}
}
