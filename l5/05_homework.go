package main

import "fmt"

// Урок 5.5 — ДОМАШКА: две задачи с двумерными срезами.
// Описание задач — в README.md. Решения — в resolve/05_homework.go.

// === Задача 1. transpose ===
//   transpose([][]int{{1, 2, 3}, {4, 5, 6}})
//     -> [[1 4] [2 5] [3 6]]
//
//   transpose([][]int{{1, 2}, {3, 4}, {5, 6}})
//     -> [[1 3 5] [2 4 6]]
//
//   transpose([][]int{})
//     -> []
//
// func transpose(matrix [][]int) [][]int {
// 	return nil
// }

// === Задача 2. spiralOrder ===
//   spiralOrder([][]int{
//       {1, 2, 3},
//       {4, 5, 6},
//       {7, 8, 9},
//   })
//     -> [1 2 3 6 9 8 7 4 5]
//
//   spiralOrder([][]int{
//       {1, 2, 3, 4},
//       {5, 6, 7, 8},
//       {9, 10, 11, 12},
//   })
//     -> [1 2 3 4 8 12 11 10 9 5 6 7]
//
//   spiralOrder([][]int{{1, 2, 3, 4, 5}})
//     -> [1 2 3 4 5]
//
//   spiralOrder([][]int{{1}, {2}, {3}})
//     -> [1 2 3]
//
//   spiralOrder([][]int{})
//     -> []
//
// func spiralOrder(matrix [][]int) []int {
// 	return nil
// }

func homework() {
	fmt.Println("Домашка: 2D-задачи")

	// --- вызовы для проверки ---

	// fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))         // [[1 4] [2 5] [3 6]]
	// fmt.Println(transpose([][]int{{1, 2}, {3, 4}, {5, 6}}))       // [[1 3 5] [2 4 6]]
	// fmt.Println(transpose([][]int{}))                              // []

	// fmt.Println(spiralOrder([][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// })) // [1 2 3 6 9 8 7 4 5]

	// fmt.Println(spiralOrder([][]int{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// })) // [1 2 3 4 8 12 11 10 9 5 6 7]

	// fmt.Println(spiralOrder([][]int{{1, 2, 3, 4, 5}})) // [1 2 3 4 5]
	// fmt.Println(spiralOrder([][]int{{1}, {2}, {3}}))   // [1 2 3]
	// fmt.Println(spiralOrder([][]int{}))                 // []
}
