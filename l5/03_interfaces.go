package main

import (
	"fmt"
)

// Урок 5.3 — ДЕМО: видимость и интерфейсы.
// Теория — в README.md.

// --- инкапсуляция через видимость на уровне пакета ---

type User struct {
	Name     string
	password string // приватное: только методы этого пакета имеют к нему доступ
}

// SetPassword — публичный сеттер. Здесь мог бы быть хеш, валидация, логика «пароль не короче 8 символов»;
// для демо оставляем простое присваивание плюс одно правило.
func (u *User) SetPassword(p string) error {
	if len(p) < 4 {
		return fmt.Errorf("пароль слишком короткий: %d символов", len(p))
	}
	u.password = "hashed:" + p // фейковый хеш для иллюстрации
	return nil
}

// CheckPassword — публичный геттер по сути, но он не отдаёт сырое значение наружу,
// а сам сравнивает то, что лежит внутри, с тем, что пришло извне.
func (u *User) CheckPassword(p string) bool {
	return u.password == "hashed:"+p
}

func encapsulationDemo() {
	u := &User{Name: "nikita"}

	if err := u.SetPassword("123"); err != nil {
		fmt.Println("ошибка установки пароля:", err)
	}
	if err := u.SetPassword("secret"); err != nil {
		fmt.Println("ошибка установки пароля:", err)
	} else {
		fmt.Println("пароль установлен для пользователя:", u.Name)
	}

	fmt.Println("проверка верного пароля:   ", u.CheckPassword("secret"))
	fmt.Println("проверка неверного пароля: ", u.CheckPassword("wrong"))

}

// --- интерфейс ---

// Animal — это контракт: всё, что умеет Speak() и Name(), считается Animal.
// Никаких полей. Никаких реализаций. Только перечень методов.
type Animal interface {
	Name() string
	Speak() string
}

// --- реализация: Dog ---
type Dog struct {
	nick string
}

func (d Dog) Name() string  { return d.nick }
func (d Dog) Speak() string { return "Гав-гав!" }

// --- реализация: Cat ---
type Cat struct {
	nick string
	mood string // "happy" или "grumpy"
}

func (c Cat) Name() string { return c.nick }
func (c Cat) Speak() string {
	if c.mood == "grumpy" {
		return "Ш-ш-ш..."
	}
	return "Мяу!"
}

// --- реализация: Cow ---
type Cow struct {
	nick string
}

func (c Cow) Name() string  { return c.nick }
func (c Cow) Speak() string { return "Му-у-у" }

// describe принимает АБСТРАКТНОЕ поведение — Animal. Внутри ему всё равно,
// корова там, собака или кошка: он работает через методы интерфейса.
func describe(a Animal) string {
	return fmt.Sprintf("%s говорит: %s", a.Name(), a.Speak())
}

func interfacesDemo() {
	fmt.Println("=== инкапсуляция ===")
	encapsulationDemo()

	fmt.Println("\n=== интерфейс Animal ===")

	// Срез разных конкретных типов под одним интерфейсом — классический полиморфизм.
	animals := []Animal{
		Dog{nick: "Бобик"},
		Cat{nick: "Мурка", mood: "happy"},
		Cat{nick: "Барсик", mood: "grumpy"},
		Cow{nick: "Зорька"},
	}

	for _, a := range animals {
		fmt.Println(describe(a))
	}
}
