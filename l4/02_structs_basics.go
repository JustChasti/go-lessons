package main

import "fmt"

// --- маленькая структура для базового показа ---
type Point struct {
	X, Y int
}

// --- структура, которую будем МЕНЯТЬ изнутри функции ---
type Person struct {
	Name string
	Age  int
	City string
}

// --- БОЛЬШАЯ структура, которую дорого копировать при передаче ---
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Timeout  int
	MaxConns int
	LogLevel string
	Debug    bool
}

// --- встраивание ---
type Employee struct {
	Person
	Company string
	Salary  int
}

// =====================================================
// Случай 1: функция должна ИЗМЕНИТЬ структуру
// =====================================================

// Принимает Person ПО ЗНАЧЕНИЮ: получает копию. Снаружи изменений не будет.
func birthdayValue(p Person) {
	p.Age++ // меняется локальная копия
}

// Принимает УКАЗАТЕЛЬ: получает адрес оригинала. Изменения видны снаружи.
func birthdayPointer(p *Person) {
	p.Age++ // эквивалент (*p).Age++
}

// =====================================================
// Случай 2: функция читает БОЛЬШУЮ структуру
// Передаём указатель, чтобы не копировать её целиком.
// =====================================================

func printConfig(c *Config) {
	fmt.Printf("  %s@%s:%d/%s\n", c.Username, c.Host, c.Port, c.Database)
	fmt.Printf("  timeout=%ds, maxConns=%d, logLevel=%s, debug=%v\n",
		c.Timeout, c.MaxConns, c.LogLevel, c.Debug)
}

func structsBasics() {
	// === создание структур ===

	// 1. нулевое значение — все поля в своих нулях
	var p1 Point
	fmt.Printf("нулевое:   %v,  тип %T\n", p1, p1)

	// 2. литерал со значениями — основной способ
	p2 := Point{X: 3, Y: 4}
	fmt.Printf("значение:  %v,  тип %T\n", p2, p2)

	// 3. литерал-указатель — когда сразу нужен *T
	p3 := &Point{X: 5, Y: 6}
	fmt.Printf("указатель: %v, тип %T (амперсанд в выводе — признак указателя)\n", p3, p3)

	// доступ к полю работает одинаково для значения и указателя:
	// Go сам разыменовывает при p3.X
	fmt.Println("p2.X =", p2.X, " p3.X =", p3.X)

	// =====================================================
	// Случай 1: меняем структуру изнутри функции
	// =====================================================
	user := Person{Name: "Никита", Age: 25, City: "Москва"}
	fmt.Println("\n--- изменение изнутри функции ---")
	fmt.Println("до:", user.Age)

	birthdayValue(user) // передали КОПИЮ
	fmt.Println("после birthdayValue(user):   ", user.Age, "(не изменилось)")

	birthdayPointer(&user) // передали АДРЕС
	fmt.Println("после birthdayPointer(&user):", user.Age, "(изменилось)")

	// =====================================================
	// Случай 2: большая структура, передаём указатель
	// чтобы не копировать каждый раз
	// =====================================================
	conf := Config{
		Host:     "localhost",
		Port:     5432,
		Username: "admin",
		Password: "secret",
		Database: "shop",
		Timeout:  30,
		MaxConns: 100,
		LogLevel: "info",
		Debug:    true,
	}

	fmt.Println("\n--- большая структура по указателю ---")
	printConfig(&conf) // передаём АДРЕС, а не всю структуру

	// =====================================================
	// Встраивание
	// =====================================================
	emp := Employee{
		Person:  Person{Name: "Аня", Age: 30, City: "Спб"},
		Company: "Acme",
		Salary:  100_000,
	}
	fmt.Println("\n--- встраивание ---")
	fmt.Println("emp.Name (через встраивание):", emp.Name)
	fmt.Println("emp.Person.Name (явный путь):", emp.Person.Name)
	fmt.Printf("вся структура: %+v\n", emp)
}
