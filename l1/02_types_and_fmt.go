package main

import "fmt"

// Урок 1.2 — Преобразование типов и форматированный вывод
// Запуск: вставь тело функции main2() в main в файл main.go
// и запусти: go run main.go

func main2() {

	// === Преобразование типов ===
	// В Go нет автоматического преобразования — всё явно

	var intNum int = 42
	var floatNum float64 = float64(intNum) // int -> float64
	var backToInt int = int(floatNum)      // float64 -> int (дробная часть отрезается!)

	fmt.Println("int:", intNum)
	fmt.Println("float64:", floatNum)
	fmt.Println("обратно int:", backToInt)

	// Осторожно с дробными!
	var precise float64 = 3.99
	var truncated int = int(precise) // будет 3, не 4!
	fmt.Println("\n3.99 в int =", truncated, "(не округляет, а отрезает)")

	// Числа разных типов нельзя складывать напрямую
	var a int32 = 10
	var b int64 = 20
	// var c = a + b  // <- ОШИБКА! разные типы
	var c = int64(a) + b // <- так правильно
	fmt.Println("\nint32 + int64 =", c)

	// === Форматированный вывод (fmt.Printf) ===

	name := "Владимир"
	age := 25
	height := 1.78
	isHappy := true

	fmt.Println("\n--- Printf ---")

	// %s — строка
	fmt.Printf("Имя: %s\n", name)

	// %d — целое число
	fmt.Printf("Возраст: %d лет\n", age)

	// %f — дробное число, %.2f — два знака после запятой
	fmt.Printf("Рост: %.2f м\n", height)

	// %t — булевое значение
	fmt.Printf("Счастлив: %t\n", isHappy)

	// %v — универсальный (Go сам подберёт формат)
	fmt.Printf("Всё вместе: %v, %v, %v\n", name, age, height)
	// почему не %v для всех? — потому что так удобнее контролировать - выдаст ошибку если что-то не так, а %v может скрыть проблему

	// %T — показать тип переменной
	fmt.Printf("\n--- Типы ---\n")
	fmt.Printf("name: %T\n", name)
	fmt.Printf("age: %T\n", age)
	fmt.Printf("height: %T\n", height)
	fmt.Printf("isHappy: %T\n", isHappy)
}
