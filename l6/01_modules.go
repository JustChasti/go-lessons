package main

import (
	"fmt"

	"l6/geom"
)

// Урок 6.1 — ДЕМО: модули, подпакеты и видимость.
// Теория — в README.md.
//
// Подпакет лежит в l6/geom. Импортируется по пути "l6/geom" — это
// имя модуля из go.mod ("l6") плюс путь к подпапке внутри модуля.

func modulesDemo() {
	// Используем экспортированные типы из подпакета: имя пакета — префикс.
	shapes := []geom.Shape{
		geom.Circle{Radius: 3},
		geom.Rectangle{Width: 4, Height: 5},
	}

	// Экспортированную функцию SumAreas вызываем без проблем.
	total := geom.SumAreas(shapes)
	fmt.Printf("сумма площадей: %.2f\n", total)

	// А вот это раскомментировать НЕ получится — компилятор скажет
	// "cannot refer to unexported name geom.roundTo":
	//
	// _ = geom.roundTo(1.234, 2)
	//
	// roundTo живёт внутри пакета geom (имя со строчной буквы),
	// и снаружи его не видно. Это и есть инкапсуляция в Go-смысле:
	// граница приватности — пакет, а не структура.
}
