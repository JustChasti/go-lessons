package geom

import "math"

// roundTo — округляет число до digits знаков после запятой.
// Имя начинается с маленькой буквы, поэтому функция видна
// только внутри пакета geom. Из main.go её не вызвать.
func roundTo(value float64, digits int) float64 {
	factor := math.Pow(10, float64(digits))
	return math.Round(value*factor) / factor
}
