package main

import "fmt"

// Урок 5.1 — ПРАКТИКА: обработка срезов через циклы и рекурсию.
// Описание задач — в README.md.

// === Задача 1. reverseSlice ===
//   reverseSlice([]int{1, 2, 3, 4, 5}) -> [5 4 3 2 1]
//   reverseSlice([]int{42})            -> [42]
//   reverseSlice([]int{})              -> []
//
// func reverseSlice(nums []int) []int {
// 	return nil
// }

// === Задача 2. filterEvens ===
//   filterEvens([]int{1, 2, 3, 4, 5, 6}) -> [2 4 6]
//   filterEvens([]int{1, 3, 5})          -> []
//   filterEvens([]int{})                 -> []
//
// func filterEvens(nums []int) []int {
// 	return nil
// }

// === Задача 3. removeAt ===
//   removeAt([]int{10, 20, 30, 40}, 1)  -> [10 30 40]
//   removeAt([]int{10, 20, 30, 40}, 0)  -> [20 30 40]
//   removeAt([]int{10, 20, 30, 40}, 5)  -> [10 20 30 40]
//   removeAt([]int{10, 20, 30, 40}, -1) -> [10 20 30 40]
//
// func removeAt(nums []int, idx int) []int {
// 	return nil
// }

// === Задача 4. dedupSorted ===
//   dedupSorted([]int{1, 1, 2, 3, 3, 3, 4}) -> [1 2 3 4]
//   dedupSorted([]int{1, 2, 3})             -> [1 2 3]
//   dedupSorted([]int{7, 7, 7, 7})          -> [7]
//   dedupSorted([]int{})                    -> []
//
// func dedupSorted(nums []int) []int {
// 	return nil
// }

// === Задача 5. sumRec ===
//   sumRec([]int{1, 2, 3, 4, 5}) -> 15
//   sumRec([]int{-1, 1})         -> 0
//   sumRec([]int{})              -> 0
//
// func sumRec(nums []int) int {
// 	return 0
// }

// === Задача 6. maxRec ===
//   maxRec([]int{3, 1, 4, 1, 5, 9, 2, 6}) -> 9
//   maxRec([]int{-7})                     -> -7
//   maxRec([]int{})                       -> 0
//
// func maxRec(nums []int) int {
// 	return 0
// }

func slicesPractice() {
	fmt.Println("Практика: срезы и циклы")

	// --- вызовы для проверки ---

	// fmt.Println(reverseSlice([]int{1, 2, 3, 4, 5})) // [5 4 3 2 1]
	// fmt.Println(reverseSlice([]int{42}))            // [42]
	// fmt.Println(reverseSlice([]int{}))              // []

	// fmt.Println(filterEvens([]int{1, 2, 3, 4, 5, 6})) // [2 4 6]
	// fmt.Println(filterEvens([]int{1, 3, 5}))          // []
	// fmt.Println(filterEvens([]int{}))                 // []

	// fmt.Println(removeAt([]int{10, 20, 30, 40}, 1))  // [10 30 40]
	// fmt.Println(removeAt([]int{10, 20, 30, 40}, 0))  // [20 30 40]
	// fmt.Println(removeAt([]int{10, 20, 30, 40}, 5))  // [10 20 30 40]
	// fmt.Println(removeAt([]int{10, 20, 30, 40}, -1)) // [10 20 30 40]

	// fmt.Println(dedupSorted([]int{1, 1, 2, 3, 3, 3, 4})) // [1 2 3 4]
	// fmt.Println(dedupSorted([]int{1, 2, 3}))             // [1 2 3]
	// fmt.Println(dedupSorted([]int{7, 7, 7, 7}))          // [7]
	// fmt.Println(dedupSorted([]int{}))                    // []

	// fmt.Println(sumRec([]int{1, 2, 3, 4, 5})) // 15
	// fmt.Println(sumRec([]int{-1, 1}))         // 0
	// fmt.Println(sumRec([]int{}))              // 0

	// fmt.Println(maxRec([]int{3, 1, 4, 1, 5, 9, 2, 6})) // 9
	// fmt.Println(maxRec([]int{-7}))                     // -7
	// fmt.Println(maxRec([]int{}))                       // 0
}
