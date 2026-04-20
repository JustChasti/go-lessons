package main

import "fmt"

// Урок 3.3 — РЕШЕНИЯ: рекурсия.

// power — построчный перевод определения: base^0 = 1, base^n = base * base^(n-1).
// Линейная глубина рекурсии — для небольших степеней этого достаточно.
func power(base, exp int) int {
	if exp == 0 {
		return 1
	}
	return base * power(base, exp-1)
}

// gcd — алгоритм Евклида. База: b == 0. На каждом шаге аргументы уменьшаются.
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// sumDigitsRec — отрицательность обрабатываем ДО рекурсии
func sumDigitsRec(n int) int {
	if n < 0 {
		n = -n
	}
	if n < 10 {
		return n
	}
	return (n % 10) + sumDigitsRec(n/10)
}

// reverseRec — берём первый символ и приклеиваем в конец перевёрнутого хвоста.
// []rune используется чтобы корректно работать с юникодом.
func reverseRec(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return ""
	}
	return reverseRec(string(r[1:])) + string(r[0])
}

// countPaths — наивная рекурсия по определению:
// из клетки (r, c) путь ведёт либо через (r-1, c), либо через (r, c-1).
// База — когда одна из сторон стала 1: путь по краю единственный.
func countPaths(rows, cols int) int {
	if rows == 1 || cols == 1 {
		return 1
	}
	return countPaths(rows-1, cols) + countPaths(rows, cols-1)
}

func recursions() {
	fmt.Println("=== Рекурсия ===")

	fmt.Println(power(2, 0))  // 1
	fmt.Println(power(2, 10)) // 1024
	fmt.Println(power(3, 4))  // 81

	fmt.Println(gcd(18, 48))  // 6
	fmt.Println(gcd(100, 75)) // 25
	fmt.Println(gcd(13, 7))   // 1
	fmt.Println(gcd(10, 0))   // 10

	fmt.Println(sumDigitsRec(1234)) // 10
	fmt.Println(sumDigitsRec(0))    // 0
	fmt.Println(sumDigitsRec(-99))  // 18

	fmt.Println(reverseRec("hello")) // olleh
	fmt.Println(reverseRec("Go!"))   // !oG
	fmt.Println(reverseRec(""))      //

	fmt.Println(countPaths(1, 1)) // 1
	fmt.Println(countPaths(2, 2)) // 2
	fmt.Println(countPaths(3, 3)) // 6
	fmt.Println(countPaths(4, 4)) // 20
}
