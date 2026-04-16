package main

import "fmt"

// Урок 2.4 — ПРАКТИКА
//
// Как пользоваться:
//   1. Раскомментируй нужное задание (функцию и описание).
//   2. Допиши реализацию вместо заглушки.
//   3. В main4() раскомментируй вызов, запусти main4() из main.go.

// === Задание 1. isLeapYear ===
// Год високосный, если он делится на 4, НО не делится на 100,
// кроме случаев, когда делится на 400.
//   2000 -> true  (делится на 400)
//   1900 -> false (делится на 100, но не на 400)
//   2024 -> true
//   2023 -> false
//
// func isLeapYear(year int) bool {
// 	return false
// }

// === Задание 2. fizzBuzz ===
// Напечатай числа от 1 до n включительно, по одному на строку.
// Вместо кратных 3  — "Fizz".
// Вместо кратных 5  — "Buzz".
// Вместо кратных 15 — "FizzBuzz".
//   fizzBuzz(15) -> 1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz
//
// func fizzBuzz(n int) {
// }

// === Задание 3. maxInSlice ===
// Верни максимальное число в срезе. Для пустого среза верни 0.
//   maxInSlice([]int{3, 1, 4, 1, 5, 9, 2, 6}) -> 9
//   maxInSlice([]int{-5, -2, -9})             -> -2
//   maxInSlice([]int{})                       -> 0
//
// func maxInSlice(nums []int) int {
// 	return 0
// }

// === Задание 4. isPrime ===
// Простое число делится только на 1 и на само себя.
//   0, 1 и отрицательные — НЕ простые.
//   2 — простое.
//   4 — нет, 7 — да, 9 — нет, 17 — да.
// Подсказка: достаточно проверить делители до sqrt(n).
// Ещё подсказка: sqrt не обязателен — хватит условия i*i <= n.
//
// func isPrime(n int) bool {
// 	return false
// }

// === Задание 5. reverseSlice ===
// Верни НОВЫЙ срез с элементами в обратном порядке.
// Исходный срез менять не надо.
//   reverseSlice([]int{1, 2, 3, 4}) -> []int{4, 3, 2, 1}
//   reverseSlice([]int{})           -> []int{}
//
// func reverseSlice(nums []int) []int {
// 	return nil
// }

// === Задание 6. digitSum ===
// Сумма цифр числа. Отрицательные считай как положительные.
//   digitSum(1234) -> 10
//   digitSum(-99)  -> 18
//   digitSum(0)    -> 0
//
// func digitSum(n int) int {
// 	return 0
// }

// === Задание 7. countVowels ===
// Сосчитай гласные (a, e, i, o, u) в строке. Регистр игнорируй.
// Считаем только английские гласные.
//   countVowels("Hello, World!") -> 3
//   countVowels("Go")            -> 1
//   countVowels("bcd")           -> 0
// Подсказка: for range по строке даёт руны, а не байты.
//
// func countVowels(s string) int {
// 	return 0
// }

func main4() {
	fmt.Println("Практика: раскомментируй задание и добавь вызов ниже.")

	// --- вызовы для проверки (раскомментируй, когда функция готова) ---

	// fmt.Println(isLeapYear(2000)) // true
	// fmt.Println(isLeapYear(1900)) // false
	// fmt.Println(isLeapYear(2024)) // true

	// fizzBuzz(20)

	// fmt.Println(maxInSlice([]int{3, 1, 4, 1, 5, 9, 2, 6})) // 9
	// fmt.Println(maxInSlice([]int{-5, -2, -9}))             // -2
	// fmt.Println(maxInSlice([]int{}))                       // 0

	// fmt.Println(isPrime(2))  // true
	// fmt.Println(isPrime(4))  // false
	// fmt.Println(isPrime(17)) // true

	// fmt.Println(reverseSlice([]int{1, 2, 3, 4})) // [4 3 2 1]

	// fmt.Println(digitSum(1234)) // 10
	// fmt.Println(digitSum(-99))  // 18

	// fmt.Println(countVowels("Hello, World!")) // 3
	// fmt.Println(countVowels("Go"))            // 1
}
