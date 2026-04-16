package main

import "fmt"

// Урок 2.1 — Условия: if / else, switch

func main1() {

	// === if / else ===
	age := 18
	if age >= 18 {
		fmt.Println("совершеннолетний")
	} else {
		fmt.Println("несовершеннолетний")
	}

	// === if / else if / else ===
	score := 75
	if score >= 90 {
		fmt.Println("отлично")
	} else if score >= 70 {
		fmt.Println("хорошо")
	} else if score >= 50 {
		fmt.Println("удовлетворительно")
	} else {
		fmt.Println("неудовлетворительно")
	}

	// === Короткая форма: объявление переменной прямо в if ===
	// n существует только внутри if/else
	if n := 42; n%2 == 0 {
		fmt.Println(n, "чётное")
	} else {
		fmt.Println(n, "нечётное")
	}
	// здесь n уже недоступно
	// fmt.Println(n)
	// n, _ := divide(6, 3) пример с функцией

	// === switch по значению ===
	day := 4
	switch day {
	case 1:
		fmt.Println("понедельник")
	case 2:
		fmt.Println("вторник")
	case 3:
		fmt.Println("среда")
	case 6, 7:
		// несколько значений — через запятую
		fmt.Println("выходной")
	default:
		fmt.Println("другой день")
	}

	// === switch без выражения (замена длинной цепочки if / else if) ===
	temp := 15
	switch {
	case temp < 0:
		fmt.Println("холодно")
	case temp < 20:
		fmt.Println("прохладно")
	case temp < 30:
		fmt.Println("тепло")
	default:
		fmt.Println("жарко")
	}

}
