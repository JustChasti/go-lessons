package main

import "fmt"

// Урок 5.1 — РЕШЕНИЯ: обработка срезов через циклы и рекурсию.

// reverseSlice — заранее выделяем результат нужной длины через make
// и заполняем его одним проходом, идя по исходному с конца.
// Исходный срез не трогаем.
func reverseSlice(nums []int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[len(nums)-1-i] = v
	}
	return result
}

// filterEvens — классическая фильтрация. Стартуем с пустого результата
// и через append добавляем подходящие элементы.
// Для пустого входа вернём пустой срез (не nil), как договорились в README.
func filterEvens(nums []int) []int {
	result := make([]int, 0, len(nums))
	for _, v := range nums {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

// removeAt — собираем результат из двух частей вокруг idx.
// Делаем копию, чтобы не трогать внутренний массив исходного среза.
func removeAt(nums []int, idx int) []int {
	if idx < 0 || idx >= len(nums) {
		// Индекс невалиден — возвращаем копию как есть.
		copyNums := make([]int, len(nums))
		copy(copyNums, nums)
		return copyNums
	}
	result := make([]int, 0, len(nums)-1)
	result = append(result, nums[:idx]...)
	result = append(result, nums[idx+1:]...)
	return result
}

// dedupSorted — два указателя.
// i идёт по входу, w — позиция куда писать в результате.
// Поскольку срез отсортирован, дубликаты идут подряд: достаточно сравнивать
// текущий элемент с последним записанным.
func dedupSorted(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	result := make([]int, 0, len(nums))
	result = append(result, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] != result[len(result)-1] {
			result = append(result, nums[i])
		}
	}
	return result
}

// sumRec — прямая рекурсия по определению:
// сумма пустого среза 0; сумма непустого — голова + сумма хвоста.
// nums[1:] — это новый заголовок над тем же массивом, без копирования элементов.
func sumRec(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + sumRec(nums[1:])
}

// maxRec — та же схема «голова + хвост», только вместо сложения выбираем максимум.
// База: один элемент — он и есть максимум.
// Особый случай: пустой срез — договорились возвращать 0.
func maxRec(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	tailMax := maxRec(nums[1:])
	if nums[0] > tailMax {
		return nums[0]
	}
	return tailMax
}

func slicesPractice() {
	fmt.Println("=== Срезы и циклы ===")

	fmt.Println(reverseSlice([]int{1, 2, 3, 4, 5})) // [5 4 3 2 1]
	fmt.Println(reverseSlice([]int{42}))            // [42]
	fmt.Println(reverseSlice([]int{}))              // []

	fmt.Println(filterEvens([]int{1, 2, 3, 4, 5, 6})) // [2 4 6]
	fmt.Println(filterEvens([]int{1, 3, 5}))          // []
	fmt.Println(filterEvens([]int{}))                 // []

	fmt.Println(removeAt([]int{10, 20, 30, 40}, 1))  // [10 30 40]
	fmt.Println(removeAt([]int{10, 20, 30, 40}, 0))  // [20 30 40]
	fmt.Println(removeAt([]int{10, 20, 30, 40}, 5))  // [10 20 30 40]
	fmt.Println(removeAt([]int{10, 20, 30, 40}, -1)) // [10 20 30 40]

	fmt.Println(dedupSorted([]int{1, 1, 2, 3, 3, 3, 4})) // [1 2 3 4]
	fmt.Println(dedupSorted([]int{1, 2, 3}))             // [1 2 3]
	fmt.Println(dedupSorted([]int{7, 7, 7, 7}))          // [7]
	fmt.Println(dedupSorted([]int{}))                    // []

	fmt.Println(sumRec([]int{1, 2, 3, 4, 5})) // 15
	fmt.Println(sumRec([]int{-1, 1}))         // 0
	fmt.Println(sumRec([]int{}))              // 0

	fmt.Println(maxRec([]int{3, 1, 4, 1, 5, 9, 2, 6})) // 9
	fmt.Println(maxRec([]int{-7}))                     // -7
	fmt.Println(maxRec([]int{}))                       // 0
}
