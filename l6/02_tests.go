package main

import (
	"fmt"
	"strings"
)

// Урок 6.2 — ДЕМО: что мы тестируем.
// Сами тесты — в файле 02_tests_test.go (он попадает в сборку только при go test).
//
// Здесь — две маленькие чистые функции. Чистая функция — та, у которой
// результат однозначно определяется аргументами и нет побочных эффектов.
// Такие функции тестировать проще всего: подаёшь вход — проверяешь выход.

// normalizeName — обрезает пробелы по краям и приводит к нижнему регистру.
// Удобно для сравнения логинов: " Alice " и "alice" должны считаться одинаковыми.
func normalizeName(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// abs — возвращает модуль числа.
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// testsDemo — запускается из main.go как обычная функция.
// Сами тесты запускаются командой `go test .` из папки l6 — она найдёт
// все *_test.go рядом и выполнит функции, начинающиеся с Test*.
func testsDemo() {
	fmt.Println("normalizeName(\"  Alice  \") =", normalizeName("  Alice  "))
	fmt.Println("abs(-7) =", abs(-7))
	fmt.Println("чтобы запустить тесты — выполни в этой папке: go test .")
	fmt.Println("чтобы увидеть детали каждой проверки — go test -v .")
}
