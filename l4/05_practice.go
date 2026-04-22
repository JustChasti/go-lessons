package main

import "fmt"

// Урок 4.5 — ПРАКТИКА
//
// Порядок работы:
//   1. Ознакомься с выдачей браузера по адресу https://api.github.com/users/justchasti (это мой гитхаб).
// это JSON-объект с кучей полей, но нам нужны только некоторые из них, обрати внимание на их структуру и типы
//   2. В main.go раскомментируй practice() и закомментируй остальные.
//   3. Реализуй задачи по очереди, раскомментируя нужные блоки ниже.
//   4. Запусти `go run .`.

// === Задача 1. Структура GitHubUser ===
// Публичный эндпоинт: https://api.github.com/users/{username}
// (без ключа; лимит 60 запросов в час с одного IP — для учебного примера хватит).
// Минимум, который тебе понадобится:
//   login         string
//   id            int
//   name          string
//   bio           string
//   public_repos  int
//   followers     int
//   following     int
//   created_at    string
// Не забудь json-теги.
//
// type GitHubUser struct {
// }

// === Задача 2. fetchGitHubUser ===
// Сделать GET на https://api.github.com/users/{username},
// прочитать тело, распарсить в GitHubUser.
// Не забудь: таймаут, проверка статуса, defer resp.Body.Close().
//
// func fetchGitHubUser(username string) (GitHubUser, error) {
// 	return GitHubUser{}, nil
// }

// === Задача 3. saveUsers ===
// Сохранить срез юзеров в JSON-файл с отступами (MarshalIndent + os.WriteFile).
//
// func saveUsers(users []GitHubUser, path string) error {
// 	return nil
// }

func practice() {
	fmt.Println("Практика: реализуй задачи в 05_practice.go")
}
