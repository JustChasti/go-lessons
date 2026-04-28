package geom

import "math"

// Урок 6.1 — подпакет geom.
// Имена с большой буквы (Shape, Circle, Rectangle, SumAreas) экспортируются —
// их видно из других пакетов через импорт "l6/geom".
// Имена с маленькой буквы (roundTo в round.go) — внутренние, снаружи их нет.

// Shape — общий контракт: всё, что умеет считать свою площадь.
type Shape interface {
	Area() float64
}

// Circle — круг с радиусом.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle — прямоугольник со сторонами Width и Height.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// SumAreas — публичная функция, считает суммарную площадь набора фигур.
// Внутри использует приватный helper roundTo — снаружи к нему не доберёшься.
func SumAreas(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return roundTo(total, 2)
}
