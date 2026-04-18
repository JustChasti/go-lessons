package main

import "fmt"

// Урок 3.1 — РЕШЕНИЯ: условия.

// ticketPrice — возраст приоритетнее статуса. Сначала отсекаем невалидный вход,
// потом детские случаи, и только для взрослых проверяем флаг студента.
func ticketPrice(age int, isStudent bool) int {
	if age < 0 {
		return -1
	}
	if age <= 5 {
		return 0
	}
	if age <= 17 {
		return 300
	}
	if isStudent {
		return 400
	}
	if age <= 64 {
		return 800
	}
	return 500
}

// classifyTriangle — две стадии: сначала валидация, потом классификация.
// Если их смешать — логика станет непрозрачной.
func classifyTriangle(a, b, c int) string {
	if a <= 0 || b <= 0 || c <= 0 {
		return "невалидный"
	}
	if a+b <= c || a+c <= b || b+c <= a {
		return "невалидный"
	}
	switch {
	case a == b && b == c:
		return "равносторонний"
	case a == b || b == c || a == c:
		return "равнобедренный"
	default:
		return "разносторонний"
	}
}

// bmiCategory — кусочная функция.
func bmiCategory(weightKg, heightM float64) string {
	if weightKg <= 0 || heightM <= 0 {
		return "некорректно"
	}
	bmi := weightKg / (heightM * heightM)
	switch {
	case bmi < 18.5:
		return "недобор"
	case bmi < 25:
		return "норма"
	case bmi < 30:
		return "лишний вес"
	default:
		return "ожирение"
	}
}

// nextLight — классический switch по значению.
func nextLight(current string) string {
	switch current {
	case "red":
		return "green"
	case "green":
		return "yellow"
	case "yellow":
		return "red"
	default:
		return "error"
	}
}

// dayKind — switch без выражения позволяет использовать условия в case.
func dayKind(day int) string {
	switch {
	case day >= 1 && day <= 5:
		return "рабочий"
	case day == 6 || day == 7:
		return "выходной"
	default:
		return "некорректно"
	}
}

func conditions() {
	fmt.Println("=== Условия ===")

	fmt.Println(ticketPrice(4, false))  // 0
	fmt.Println(ticketPrice(10, true))  // 300
	fmt.Println(ticketPrice(20, true))  // 400
	fmt.Println(ticketPrice(30, false)) // 800
	fmt.Println(ticketPrice(70, false)) // 500
	fmt.Println(ticketPrice(-3, false)) // -1

	fmt.Println(classifyTriangle(3, 4, 5))  // разносторонний
	fmt.Println(classifyTriangle(5, 5, 8))  // равнобедренный
	fmt.Println(classifyTriangle(7, 7, 7))  // равносторонний
	fmt.Println(classifyTriangle(1, 2, 10)) // невалидный
	fmt.Println(classifyTriangle(0, 4, 5))  // невалидный

	fmt.Println(bmiCategory(50, 1.70))  // недобор
	fmt.Println(bmiCategory(68, 1.75))  // норма
	fmt.Println(bmiCategory(82, 1.75))  // лишний вес
	fmt.Println(bmiCategory(100, 1.70)) // ожирение
	fmt.Println(bmiCategory(0, 1.70))   // некорректно

	fmt.Println(nextLight("red"))    // green
	fmt.Println(nextLight("green"))  // yellow
	fmt.Println(nextLight("yellow")) // red
	fmt.Println(nextLight("blue"))   // error

	fmt.Println(dayKind(1)) // рабочий
	fmt.Println(dayKind(6)) // выходной
	fmt.Println(dayKind(0)) // некорректно
}
