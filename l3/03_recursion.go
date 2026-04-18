package main

import "fmt"

// Урок 3.3 — ПРАКТИКА: рекурсия
// Описание задач — в README.md.

// === Задача 1. power ===
//   power(2, 0)  -> 1
//   power(2, 10) -> 1024
//   power(3, 4)  -> 81
//
// func power(base, exp int) int {
// 	return 0
// }

// === Задача 2. gcd ===
//   gcd(48, 18)  -> 6
//   gcd(100, 75) -> 25
//   gcd(13, 7)   -> 1
//   gcd(10, 0)   -> 10
//
// func gcd(a, b int) int {
// 	return 0
// }

// === Задача 3. sumDigitsRec ===
//   sumDigitsRec(1234) -> 10
//   sumDigitsRec(0)    -> 0
//   sumDigitsRec(-99)  -> 18
//
// func sumDigitsRec(n int) int {
// 	return 0
// }

// === Задача 4. reverseRec ===
//   reverseRec("hello") -> "olleh"
//   reverseRec("Go!")   -> "!oG"
//   reverseRec("")      -> ""
//
// func reverseRec(s string) string {
// 	return ""
// }

// === Задача 5. countPaths ===
//   countPaths(1, 1) -> 1
//   countPaths(2, 2) -> 2
//   countPaths(3, 3) -> 6
//   countPaths(4, 4) -> 20
//
// func countPaths(rows, cols int) int {
// 	return 0
// }

func recursions() {
	fmt.Println("Практика по рекурсии")

	// --- вызовы для проверки ---

	// fmt.Println(power(2, 0))  // 1
	// fmt.Println(power(2, 10)) // 1024
	// fmt.Println(power(3, 4))  // 81

	// fmt.Println(gcd(48, 18))  // 6
	// fmt.Println(gcd(100, 75)) // 25
	// fmt.Println(gcd(13, 7))   // 1
	// fmt.Println(gcd(10, 0))   // 10

	// fmt.Println(sumDigitsRec(1234)) // 10
	// fmt.Println(sumDigitsRec(0))    // 0
	// fmt.Println(sumDigitsRec(-99))  // 18

	// fmt.Println(reverseRec("hello")) // olleh
	// fmt.Println(reverseRec("Go!"))   // !oG

	// fmt.Println(countPaths(1, 1)) // 1
	// fmt.Println(countPaths(2, 2)) // 2
	// fmt.Println(countPaths(3, 3)) // 6
	// fmt.Println(countPaths(4, 4)) // 20
}
