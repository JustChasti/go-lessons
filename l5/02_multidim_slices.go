package main

import "fmt"

// Урок 5.2 — ДЕМО: многомерные срезы.
// Теория — в README.md. Здесь готовый код для запуска и чтения.

// --- объявление двух способов ---
func declareDemo() {
	// 1. Литералом — когда строки известны заранее
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("матрица литералом:", matrix)

	// 2. Через make — когда размеры вычисляются
	rows, cols := 3, 4
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	// сейчас grid — это 3x4 матрица из нулей
	grid[1][2] = 99
	fmt.Println("матрица через make:", grid)
}

// --- проход двойным циклом ---
func traverseDemo() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Сумма всех элементов
	total := 0
	for i := range matrix {
		for j := range matrix[i] {
			total += matrix[i][j]
		}
	}
	fmt.Println("сумма всех элементов:", total)

	// Максимум по всей матрице
	max := matrix[0][0]
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] > max {
				max = matrix[i][j]
			}
		}
	}
	fmt.Println("максимум:", max)
}

// --- сумма по строкам ---
func rowSumsDemo() {
	matrix := [][]int{
		{1, 2, 3},
		{10, 20, 30},
		{100, 200, 300},
	}

	rowSums := make([]int, len(matrix))
	for i := range matrix {
		sum := 0
		for j := range matrix[i] {
			sum += matrix[i][j]
		}
		rowSums[i] = sum
	}
	fmt.Println("суммы по строкам:", rowSums)
}

// --- jagged-матрица: рваные ряды разной длины ---
func jaggedDemo() {
	jagged := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
		{7, 8, 9, 10},
	}

	for i := range jagged {
		fmt.Printf("строка %d (длина %d): %v\n", i, len(jagged[i]), jagged[i])
	}
}

// --- игровая карта: '#' стена, '.' проход, '@' игрок ---
func mapDemo() {
	gameMap := [][]rune{
		[]rune("##########"),
		[]rune("#..#.....#"),
		[]rune("#.@#..#..#"),
		[]rune("#...#.#..#"),
		[]rune("##########"),
	}

	// Печать карты
	for i := range gameMap {
		fmt.Println(string(gameMap[i]))
	}

	// Подсчёт открытых клеток и поиск игрока
	openCells := 0
	playerRow, playerCol := -1, -1
	for i := range gameMap {
		for j := range gameMap[i] {
			switch gameMap[i][j] {
			case '.':
				openCells++
			case '@':
				playerRow, playerCol = i, j
				openCells++ // игрок стоит на проходимой клетке
			}
		}
	}
	fmt.Println("проходимых клеток:", openCells)
	fmt.Printf("игрок в (%d, %d)\n", playerRow, playerCol)
}

func multidimSlices() {
	fmt.Println("=== объявление ===")
	declareDemo()

	fmt.Println("\n=== обход двойным циклом ===")
	traverseDemo()

	fmt.Println("\n=== сумма по строкам ===")
	rowSumsDemo()

	fmt.Println("\n=== jagged-матрица ===")
	jaggedDemo()

	fmt.Println("\n=== игровая карта ===")
	mapDemo()
}
