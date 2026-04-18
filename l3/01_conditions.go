package main

import "fmt"

// Урок 3.1 — ПРАКТИКА: условия
// Описание задач — в README.md.

// === Задача 1. ticketPrice ===
//   ticketPrice(4, false)   -> 0
//   ticketPrice(10, true)   -> 300
//   ticketPrice(20, true)   -> 400
//   ticketPrice(30, false)  -> 800
//   ticketPrice(70, false)  -> 500
//   ticketPrice(-3, false)  -> -1
//
// func ticketPrice(age int, isStudent bool) int {
// 	return 0
// }

// === Задача 2. classifyTriangle ===
//   classifyTriangle(3, 4, 5)   -> "разносторонний"
//   classifyTriangle(5, 5, 8)   -> "равнобедренный"
//   classifyTriangle(7, 7, 7)   -> "равносторонний"
//   classifyTriangle(1, 2, 10)  -> "невалидный"
//   classifyTriangle(0, 4, 5)   -> "невалидный"
//
// func classifyTriangle(a, b, c int) string {
// 	return ""
// }

// === Задача 3. bmiCategory ===
//   bmiCategory(50, 1.70)   -> "недобор"
//   bmiCategory(68, 1.75)   -> "норма"
//   bmiCategory(82, 1.75)   -> "лишний вес"
//   bmiCategory(100, 1.70)  -> "ожирение"
//   bmiCategory(0, 1.70)    -> "некорректно"
//
// func bmiCategory(weightKg, heightM float64) string {
// 	return ""
// }

// === Задача 4. nextLight ===
//   nextLight("red")     -> "green"
//   nextLight("green")   -> "yellow"
//   nextLight("yellow")  -> "red"
//   nextLight("blue")    -> "error"
//
// func nextLight(current string) string {
// 	return ""
// }

// === Задача 5. dayKind ===
//   dayKind(1)  -> "рабочий"
//   dayKind(6)  -> "выходной"
//   dayKind(0)  -> "некорректно"
//
// func dayKind(day int) string {
// 	return ""
// }

func conditions() {
	fmt.Println("Практика по условиям")

	// --- вызовы для проверки ---

	// fmt.Println(ticketPrice(4, false))  // 0
	// fmt.Println(ticketPrice(10, true))  // 300
	// fmt.Println(ticketPrice(20, true))  // 400
	// fmt.Println(ticketPrice(30, false)) // 800
	// fmt.Println(ticketPrice(70, false)) // 500
	// fmt.Println(ticketPrice(-3, false)) // -1

	// fmt.Println(classifyTriangle(3, 4, 5))  // разносторонний
	// fmt.Println(classifyTriangle(5, 5, 8))  // равнобедренный
	// fmt.Println(classifyTriangle(7, 7, 7))  // равносторонний
	// fmt.Println(classifyTriangle(1, 2, 10)) // невалидный

	// fmt.Println(bmiCategory(50, 1.70))  // недобор
	// fmt.Println(bmiCategory(68, 1.75))  // норма
	// fmt.Println(bmiCategory(82, 1.75))  // лишний вес
	// fmt.Println(bmiCategory(100, 1.70)) // ожирение
	// fmt.Println(bmiCategory(0, 1.70))   // некорректно

	// fmt.Println(nextLight("red"))    // green
	// fmt.Println(nextLight("green"))  // yellow
	// fmt.Println(nextLight("yellow")) // red
	// fmt.Println(nextLight("blue"))   // error

	// fmt.Println(dayKind(1)) // рабочий
	// fmt.Println(dayKind(6)) // выходной
	// fmt.Println(dayKind(0)) // некорректно
}
