package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Урок 5.4 — ДЕМО: обработка ошибок.
// Теория — в README.md.

// --- 1. функция, которая может ошибиться: возвращает результат и error ---

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

func divideDemo() {
	if result, err := safeDivide(10, 2); err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	if result, err := safeDivide(10, 0); err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("10 / 0 =", result)
	}
}

// --- 2. ошибка приходит из стандартной библиотеки ---

func parseAgeDemo() {
	inputs := []string{"42", "-7", "abc", ""}

	for _, s := range inputs {
		// Atoi возвращает (int, error). Обрабатываем оба значения.
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("вход %q: не удалось распарсить — %v\n", s, err)
			continue
		}
		fmt.Printf("вход %q: распарсилось как %d\n", s, n)
	}
}

// --- 3. sentinel-ошибка + errors.Is ---

// ErrUserNotFound — заранее объявленная переменная-ошибка.
// Вызывающий код может на неё ссылаться через errors.Is.
var ErrUserNotFound = errors.New("пользователь не найден")

// fakeFindUser — имитация поиска в базе данных.
// Возвращает либо имя пользователя, либо ErrUserNotFound.
func fakeFindUser(id int) (string, error) {
	users := map[int]string{
		1: "alice",
		2: "bob",
	}
	if name, ok := users[id]; ok {
		return name, nil
	}
	return "", ErrUserNotFound
}

func sentinelDemo() {
	for _, id := range []int{1, 2, 999} {
		name, err := fakeFindUser(id)
		if err != nil {
			// Конкретно эту ошибку обрабатываем особым образом.
			if errors.Is(err, ErrUserNotFound) {
				fmt.Printf("id=%d: пропускаем, такого пользователя нет\n", id)
				continue
			}
			// Любая другая ошибка — что-то более серьёзное.
			fmt.Printf("id=%d: неожиданная ошибка: %v\n", id, err)
			continue
		}
		fmt.Printf("id=%d: найден %s\n", id, name)
	}
}

// --- 4. обёртывание ошибки через %w ---

func loadConfig(path string) error {
	_, err := os.ReadFile(path)
	if err != nil {
		// Добавляем контекст, но НЕ теряем исходную ошибку:
		// %w оборачивает её, чтобы errors.Is/As могли до неё добраться.
		return fmt.Errorf("loadConfig(%q): %w", path, err)
	}
	return nil
}

func wrapDemo() {
	err := loadConfig("/no/such/file.json")
	if err != nil {
		fmt.Println("ошибка целиком:", err)

		// Можно проверить, что под обёрткой именно «нет такого файла».
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("→ распознали os.ErrNotExist под обёрткой")
		}
	}
}

// --- 5. свой тип ошибки ---

// ValidationError — структура с дополнительной информацией.
// Чтобы тип удовлетворял интерфейсу error, нужен метод Error() string.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("поле %q: %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "отрицательный возраст"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "нереалистичный возраст"}
	}
	return nil
}

func customErrorDemo() {
	for _, age := range []int{30, -5, 200} {
		err := validateAge(age)
		if err == nil {
			fmt.Printf("age=%d: ОК\n", age)
			continue
		}

		// errors.As извлекает конкретный тип из цепочки ошибок и заполняет указатель.
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("age=%d: ошибка валидации поля %q: %s\n", age, ve.Field, ve.Message)
			continue
		}
		// Если бы под рукой была ошибка другого типа — обработали бы её здесь.
		fmt.Printf("age=%d: общая ошибка: %v\n", age, err)
	}
}

func errorsDemo() {
	fmt.Println("=== деление с проверкой ===")
	divideDemo()

	fmt.Println("\n=== парсинг строки в число ===")
	parseAgeDemo()

	fmt.Println("\n=== sentinel-ошибка через errors.Is ===")
	sentinelDemo()

	fmt.Println("\n=== обёртывание через %w ===")
	wrapDemo()

	fmt.Println("\n=== свой тип ошибки + errors.As ===")
	customErrorDemo()
}
