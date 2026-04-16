package main

import "fmt"

// Урок 2.2 — Циклы: все формы for, range, break, continue

func main2() {

	// === Классический for ===
	fmt.Println("-- классический for --")
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	// === for как while (только условие) ===
	fmt.Println("-- for как while --")
	n := 10
	for n > 0 {
		fmt.Println("n =", n)
		n -= 3
	}

	// === Бесконечный цикл + break ===
	fmt.Println("-- бесконечный + break --")
	count := 0
	for {
		if count == 3 {
			break
		}
		fmt.Println("count =", count)
		count++
	}

	// === continue: пропустить итерацию ===
	fmt.Println("-- continue --")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // пропустить чётные
		}
		fmt.Println("нечётное:", i)
	}

	// === for range по срезу ===
	fmt.Println("-- range по срезу --")
	nums := []int{10, 20, 30, 40}
	for idx, val := range nums {
		fmt.Printf("index=%d value=%d\n", idx, val)
	}

	// === for range, когда индекс не нужен ===
	fmt.Println("-- range с _ --")
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println("sum =", sum)

	// === for range только по индексу (опуская значение) ===
	fmt.Println("-- range только по индексу --")
	for i, _ := range nums {
		fmt.Println("i =", i)
	}

	// === for range по строке: идёт по рунам ===
	fmt.Println("-- range по строке --")
	for i, r := range "Go!" {
		fmt.Printf("index=%d rune=%c code=%d\n", i, r, r)
	}

	// === Вложенные циклы — таблица умножения ===
	fmt.Println("-- вложенные циклы --")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d*%d=%d  ", i, j, i*j)
		}
		fmt.Println()
	}
}
