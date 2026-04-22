package main

import "fmt"

// --- массивы ---
func arraysDemo() {
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	fmt.Println("массив:", arr, "длина:", len(arr))

	words := [3]string{"go", "rust", "zig"}
	fmt.Println("массив строк:", words, "длина:", len(words))
}

// --- срезы ---
func slicesDemo() {
	s := []int{1, 2, 3}
	fmt.Println("срез:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, 4, 5)
	fmt.Println("после append:", s, "len:", len(s), "cap:", cap(s))

	// срез можно получить из массива
	arr := [5]int{10, 20, 30, 40, 50}
	sun := arr[:] // arr[:] - срез всего массива		оборачивает массив, поэтому
	//				 sun[0] = 999  	// изменит arr[0] = 999, так как sun и arr делят один массив
	sun = make([]int, len(arr)) // теперь sun - новый срез, не связанный с arr, но с той же длиной
	copy(sun, arr[:])
	sun[0] = 999
	fmt.Println("срез из всего массива:", sun, "len:", len(sun), "cap:", cap(sun))

	sub := arr[1:4]
	fmt.Println("подсрез массива:", sub, "len:", len(sub), "cap:", cap(sub))

	// пустой срез через make: длина 0, ёмкость 10
	buf := make([]int, 0, 10)
	fmt.Println("make:", buf, "len:", len(buf), "cap:", cap(buf))
}

// --- указатели ---
func pointersDemo() {
	x := 42
	p := &x
	fmt.Println("x:", x, "адрес &x:", p, "значение *p:", *p)

	*p = 100
	fmt.Println("после *p = 100, x =", x)
}

// --- передача по значению ---
func doubleByValue(n int) {
	n = n * 2 // меняем ЛОКАЛЬНУЮ копию
}

// --- передача по указателю ---
func doubleByPointer(n *int) {
	*n = *n * 2 // меняем значение по адресу
}

// --- срез: заголовок передаётся по значению, но массив под ним — общий ---
func modifySlice(s []int) {
	s[0] = 100 // изменение видно снаружи
}

func functionsDemo() {
	n := 5
	doubleByValue(n)
	fmt.Println("после doubleByValue:", n) // 5, не изменилось

	doubleByPointer(&n)
	fmt.Println("после doubleByPointer:", n) // 10

	nums := []int{1, 2, 3}
	modifySlice(nums)
	fmt.Println("после modifySlice:", nums) // [100 2 3]
}

func slicesAndPointers() {
	fmt.Println("=== массивы ===")
	arraysDemo()

	fmt.Println("\n=== срезы ===")
	slicesDemo()

	fmt.Println("\n=== указатели ===")
	pointersDemo()

	fmt.Println("\n=== передача в функции ===")
	functionsDemo()
}
