package main

import "fmt"

// Урок 2.3 — Функции: параметры, возвращаемые значения, замыкания, рекурсия

// Простая функция: два параметра, один возврат.
// Если типы соседних параметров совпадают — тип можно указать один раз.
func add(a, b int) int {
	return a + b
}

// Несколько возвращаемых значений.
// Идиома Go: (результат, ошибка). Если ошибки нет — вернуть nil.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return a / b, nil
}

// Именованные возвращаемые значения.
// quotient и remainder объявлены сразу; голый return возвращает их текущие значения.
func divmod(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

// Вариативная функция: принимает любое количество int.
// Внутри nums — обычный срез []int.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Рекурсия. База — n <= 1.
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Функция, возвращающая функцию (замыкание).
// x захвачен из внешней функции и живёт вместе с возвращённой.
func multiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func main3() {

	// === Простой вызов ===
	fmt.Println("add(3, 4) =", add(3, 4))

	// === Обработка (результат, ошибка) ===
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("10 / 3 =", result)
	}

	_, err = divide(1, 0)
	if err != nil {
		fmt.Println("ошибка:", err)
	}

	// === Именованные результаты ===
	q, r := divmod(17, 5)
	fmt.Printf("17 / 5 = %d ост %d\n", q, r)

	// === Вариативная функция ===
	fmt.Println("sum(1,2,3) =", sum(1, 2, 3))
	fmt.Println("sum(10,20,30,40) =", sum(10, 20, 30, 40))

	// Распаковка среза в вариативный аргумент через ...
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("sum(nums...) =", sum(nums...))

	// === Рекурсия ===
	fmt.Println("5! =", factorial(5))

	// === Функция как значение: присваиваем в переменную ===
	var op func(int, int) int
	op = add
	fmt.Println("op(2, 3) =", op(2, 3))

	// === Замыкание: double и triple — разные функции с разным захваченным x ===
	double := multiplier(2)
	triple := multiplier(3)
	fmt.Println("double(5) =", double(5))
	fmt.Println("triple(5) =", triple(5))

	// === Анонимная функция: можно объявить и сразу вызвать ===
	func() {
		fmt.Println("это анонимная функция, вызванная сразу")
	}()
}
