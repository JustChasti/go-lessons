package main

import "fmt"

// Урок 1.1 — Переменные и типы данных
// Запуск: вставь тело функции main1() в main в файл main.go
// и запусти: go run main.go

func main1() {

	// === Способы объявления переменных ===

	// 1. Полное объявление с типом
	var name string = "Никита"
	var age int = 30

	// 2. Тип можно не указывать — Go сам поймёт (type inference)
	var city = "Москва"

	// 3. Короткое объявление через := (только внутри функций)
	language := "Go"

	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Город:", city)
	fmt.Println("Язык:", language)

	// === Основные типы данных ===

	// Целые числа
	var counter int32 = 42     // стандартный (32 бита)
	var small int8 = 127       // от -128 до 127 (256 чисел или 2 в 8 степениб по остальным типам аналогично)
	var big int64 = 9999999999 // большие числа

	// Дробные числа
	var price float64 = 19.99
	var temperature float32 = 36.6

	// Булевый тип
	var isStudent bool = true

	// Строка
	var greeting string = "Привет!"

	// Байт и руна
	var letter byte = 'A' // один символ из таблицы ASCII (латиница, цифры, знаки)
	var emoji rune = '🚀'  // один любой символ, включая кириллицу и эмодзи

	// byte — для простых английских букв и знаков (1 байт на символ)
	// rune — для всего остального: русские буквы, иероглифы, эмодзи (до 4 байт на символ)

	fmt.Println("\n--- Типы данных ---")
	fmt.Println("counter:", counter)
	fmt.Println("small:", small)
	fmt.Println("big:", big)
	fmt.Println("price:", price)
	fmt.Println("temperature:", temperature)
	fmt.Println("isStudent:", isStudent)
	fmt.Println("greeting:", greeting)
	fmt.Println("letter:", letter, "- это числовой код буквы A")
	fmt.Println("letter как символ:", string(letter))
	fmt.Println("emoji:", string(emoji))
	fmt.Println("len(\"Привет\"):", len("Привет"), "байт, а не 6 символов!")

	// === Нулевые значения (zero values) ===
	// Если не присвоить значение — Go даст значение по умолчанию

	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string

	fmt.Println("\n--- Нулевые значения ---")
	fmt.Println("int:", defaultInt)       // 0
	fmt.Println("float64:", defaultFloat) // 0
	fmt.Println("bool:", defaultBool)     // false
	fmt.Println("string:", defaultString) // "" (пустая строка)

	// === Константы ===
	// Значение задаётся один раз и не меняется

	const pi = 3.14159
	const appName = "MyApp"

	fmt.Println("\n--- Константы ---")
	fmt.Println("pi:", pi)
	fmt.Println("appName:", appName)

	// pi = 3.14 // <- раскомментируй эту строку и увидишь ошибку
}
