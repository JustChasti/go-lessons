package main

import "fmt"

// Урок 5.5 — РЕШЕНИЯ: домашка по двумерным срезам.

// transpose — создаём новую матрицу размером N×M (где исходная M×N).
// Сначала аккуратно разбираем размеры: количество строк результата
// равно количеству столбцов исходника, и наоборот.
//
// Внутренние срезы выделяем по одному, чтобы не получить «один и тот же» массив
// под разными строками.
func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}
	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]int, cols)
	for j := 0; j < cols; j++ {
		result[j] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}

// spiralOrder — обход по спирали через четыре границы.
//
// На каждой итерации делаем четыре прохода: верх слева направо, правый край сверху вниз,
// низ справа налево, левый край снизу вверх. После каждого прохода соответствующая
// граница сдвигается «внутрь».
//
// Тонкость: после первых двух проходов матрица могла «исчерпаться» — для последних двух
// нужны проверки top <= bottom и left <= right, иначе на узких матрицах
// (одна строка или один столбец) можно пройти элементы повторно.
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	rows := len(matrix)
	cols := len(matrix[0])
	result := make([]int, 0, rows*cols)

	top, bottom := 0, rows-1
	left, right := 0, cols-1

	for top <= bottom && left <= right {
		// верх: слева направо
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		top++

		// правый край: сверху вниз
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// низ: справа налево (только если строка ещё осталась)
		if top <= bottom {
			for j := right; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
			bottom--
		}

		// левый край: снизу вверх (только если столбец ещё остался)
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

func homework() {
	fmt.Println("=== Домашка ===")

	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))   // [[1 4] [2 5] [3 6]]
	fmt.Println(transpose([][]int{{1, 2}, {3, 4}, {5, 6}})) // [[1 3 5] [2 4 6]]
	fmt.Println(transpose([][]int{}))                       // []

	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})) // [1 2 3 6 9 8 7 4 5]

	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	})) // [1 2 3 4 8 12 11 10 9 5 6 7]

	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4, 5}})) // [1 2 3 4 5]
	fmt.Println(spiralOrder([][]int{{1}, {2}, {3}}))   // [1 2 3]
	fmt.Println(spiralOrder([][]int{}))                // []
}
