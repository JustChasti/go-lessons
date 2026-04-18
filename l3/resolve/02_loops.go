package main

import (
	"fmt"
	"strings"
)

// Урок 3.2 — РЕШЕНИЯ: циклы.

// sumOfMultiples — двигаемся сразу шагом k, а не проверяем каждое число на делимость
func sumOfMultiples(limit, k int) int {
	if k <= 0 {
		return 0
	}
	total := 0
	for i := k; i <= limit; i += k {
		total += i
	}
	return total
}

// countDigits — делим на 10, пока не станет 0. Отрицательные приводим к положительным
// отдельным шагом, чтобы в цикле была только одна задача
func countDigits(n int) int {
	if n < 0 {
		n = -n
	}
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

// collatzSteps — ровно реализация определения. Выход из цикла по значению, а не по счётчику.
func collatzSteps(n int) int {
	if n <= 0 {
		return -1
	}
	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return steps
}

// isPalindromeStr — два указателя навстречу друг другу.
// ToLower нужен ради сравнения без учёта регистра, []rune — чтобы индексировать строку с возможным юникодом
func isPalindromeStr(s string) bool {
	r := []rune(strings.ToLower(s))
	i, j := 0, len(r)-1
	for i < j {
		if r[i] != r[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// minMax — инициализируем min и max первым элементом (а не MaxInt/MinInt)
func minMax(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	mn, mx := nums[0], nums[0]
	for _, v := range nums {
		if v < mn {
			mn = v
		}
		if v > mx {
			mx = v
		}
	}
	return mn, mx
}

// smallestPowerGreater — пошагово удваиваем value, пока не перерастёт limit.
// Работаем в целых числах, без float и логарифмов.
func smallestPowerGreater(limit int) int {
	if limit < 1 {
		return 0
	}
	k, value := 0, 1
	for value <= limit {
		value *= 2
		k++
	}
	return k
}

func loops() {
	fmt.Println("=== Циклы ===")

	fmt.Println(sumOfMultiples(10, 3))  // 18
	fmt.Println(sumOfMultiples(20, 5))  // 50
	fmt.Println(sumOfMultiples(100, 7)) // 735
	fmt.Println(sumOfMultiples(10, 0))  // 0

	fmt.Println(countDigits(0))    // 1
	fmt.Println(countDigits(9))    // 1
	fmt.Println(countDigits(1234)) // 4
	fmt.Println(countDigits(-42))  // 2

	fmt.Println(collatzSteps(1))  // 0
	fmt.Println(collatzSteps(6))  // 8
	fmt.Println(collatzSteps(27)) // 111
	fmt.Println(collatzSteps(0))  // -1

	fmt.Println(isPalindromeStr("level"))   // true
	fmt.Println(isPalindromeStr("RaceCar")) // true
	fmt.Println(isPalindromeStr("hello"))   // false
	fmt.Println(isPalindromeStr(""))        // true

	fmt.Println(minMax([]int{3, 1, 4, 1, 5, 9, 2, 6})) // 1 9
	fmt.Println(minMax([]int{-5}))                     // -5 -5
	fmt.Println(minMax([]int{}))                       // 0 0

	fmt.Println(smallestPowerGreater(1))     // 1
	fmt.Println(smallestPowerGreater(1000))  // 10
	fmt.Println(smallestPowerGreater(10000)) // 14
	fmt.Println(smallestPowerGreater(0))     // 0
}
