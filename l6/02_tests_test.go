package main

import "testing"

// Урок 6.2 — собственно тесты.
//
// Файлы с суффиксом _test.go компилируются только при `go test`,
// в обычной сборке `go run .` они не участвуют.
//
// Тестовая функция должна:
//   - называться с префикса Test
//   - принимать единственный аргумент *testing.T
//
// Падение теста — через t.Errorf (продолжает выполнение) или t.Fatalf
// (останавливает текущую тест-функцию). Если за весь тест ни одного
// сообщения об ошибке не было — он считается пройденным.

func TestNormalizeName(t *testing.T) {
	// Табличные тесты — стандартный приём в Go.
	// Один цикл проходит по списку случаев "вход → ожидание".
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"trim and lower", "  Alice  ", "alice"},
		{"already normal", "bob", "bob"},
		{"upper case", "BOB", "bob"},
		{"empty", "", ""},
		{"only spaces", "   ", ""},
	}

	for _, c := range cases {
		// t.Run даёт каждому случаю своё имя в выводе.
		// Если упадёт — будет видно конкретно какой.
		t.Run(c.name, func(t *testing.T) {
			got := normalizeName(c.in)
			if got != c.want {
				t.Errorf("normalizeName(%q) = %q; хотим %q", c.in, got, c.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
		{-1, 1},
	}

	for _, c := range cases {
		if got := abs(c.in); got != c.want {
			t.Errorf("abs(%d) = %d; хотим %d", c.in, got, c.want)
		}
	}
}
