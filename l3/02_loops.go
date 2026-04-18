package main

import "fmt"

// Урок 3.2 — ПРАКТИКА: циклы
// Описание задач — в README.md.

// === Задача 1. sumOfMultiples ===
//   sumOfMultiples(10, 3)  -> 18     (3+6+9)
//   sumOfMultiples(20, 5)  -> 50     (5+10+15+20)
//   sumOfMultiples(100, 7) -> 735
//   sumOfMultiples(10, 0)  -> 0
//
// func sumOfMultiples(limit, k int) int {
// 	return 0
// }

// === Задача 2. countDigits ===
//   countDigits(0)    -> 1
//   countDigits(9)    -> 1
//   countDigits(1234) -> 4
//   countDigits(-42)  -> 2
//
// func countDigits(n int) int {
// 	return 0
// }

// === Задача 3. collatzSteps ===
//   collatzSteps(1)  -> 0
//   collatzSteps(6)  -> 8
//   collatzSteps(27) -> 111
//   collatzSteps(0)  -> -1
//
// func collatzSteps(n int) int {
// 	return 0
// }

// === Задача 4. isPalindromeStr ===
//   isPalindromeStr("level")   -> true
//   isPalindromeStr("RaceCar") -> true
//   isPalindromeStr("hello")   -> false
//   isPalindromeStr("")        -> true
//
// func isPalindromeStr(s string) bool {
// 	return false
// }

// === Задача 5. minMax ===
//   minMax([]int{3, 1, 4, 1, 5, 9, 2, 6}) -> (1, 9)
//   minMax([]int{-5})                     -> (-5, -5)
//   minMax([]int{})                       -> (0, 0)
//
// func minMax(nums []int) (int, int) {
// 	return 0, 0
// }

// === Задача 6. smallestPowerGreater ===
//   smallestPowerGreater(1)     -> 1
//   smallestPowerGreater(1000)  -> 10
//   smallestPowerGreater(10000) -> 14
//   smallestPowerGreater(0)     -> 0
//
// func smallestPowerGreater(limit int) int {
// 	return 0
// }

func loops() {
	fmt.Println("Практика по циклам")

	// --- вызовы для проверки ---

	// fmt.Println(sumOfMultiples(10, 3))  // 18
	// fmt.Println(sumOfMultiples(20, 5))  // 50
	// fmt.Println(sumOfMultiples(100, 7)) // 735

	// fmt.Println(countDigits(0))    // 1
	// fmt.Println(countDigits(1234)) // 4
	// fmt.Println(countDigits(-42))  // 2

	// fmt.Println(collatzSteps(1))  // 0
	// fmt.Println(collatzSteps(6))  // 8
	// fmt.Println(collatzSteps(27)) // 111

	// fmt.Println(isPalindromeStr("level"))   // true
	// fmt.Println(isPalindromeStr("RaceCar")) // true
	// fmt.Println(isPalindromeStr("hello"))   // false

	// fmt.Println(minMax([]int{3, 1, 4, 1, 5, 9, 2, 6})) // 1 9
	// fmt.Println(minMax([]int{-5}))                     // -5 -5
	// fmt.Println(minMax([]int{}))                       // 0 0

	// fmt.Println(smallestPowerGreater(1))     // 1
	// fmt.Println(smallestPowerGreater(1000))  // 10
	// fmt.Println(smallestPowerGreater(10000)) // 14
}
