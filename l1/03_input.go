package main

import "fmt"

// Урок 1.3 — Ввод данных от пользователя
// Запуск: вставь тело функции main3() в main в файл main.go
// и запусти: go run main.go

func main3() {

	// === fmt.Scan — простой ввод ===

	var name string
	fmt.Print("Как тебя зовут? ")
	fmt.Scan(&name) // & — это указатель, пока просто запомни что он нужен для Scan

	var age int
	fmt.Print("Сколько тебе лет? ")
	fmt.Scan(&age)

	fmt.Printf("\nПривет, %s! Тебе %d лет.\n", name, age)

	// Через 10 лет
	fmt.Printf("Через 10 лет тебе будет %d лет.\n", age+10)

	// === Простая арифметика с вводом ===

	var a, b float64
	fmt.Print("\nВведи первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введи второе число: ")
	fmt.Scan(&b)

	fmt.Printf("\n--- Результаты ---\n")
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)

	if b != 0 {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
	} else {
		fmt.Println("На ноль делить нельзя!")
	}
}
