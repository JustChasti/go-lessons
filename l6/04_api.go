package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Урок 6.4 — ДЕМО: GET / POST / PUT / DELETE.
// Теория и список ресурсов — в README.md.
//
// Используем два публичных API без ключей:
//   - https://api.open-meteo.com   — настоящие данные о погоде, GET с параметрами в URL
//   - https://jsonplaceholder.typicode.com — фейковый REST на posts, поддерживает все методы

// --- структура для постов из jsonplaceholder ---
//
// id ставит сервер при создании, поэтому помечаем omitempty —
// при POST мы его не отправляем.
type Post struct {
	ID     int    `json:"id,omitempty"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Один общий клиент с таймаутом — переиспользуем во всех функциях.
// http.Get / http.Post используют DefaultClient без таймаута, что в проде опасно.
var apiClient = &http.Client{Timeout: 10 * time.Second}

// --- 1. GET с параметрами в URL ---
//
// В URL после "?" идут пары key=value, разделённые "&".
// Значения нужно url-encode-ить (пробелы превратятся в %20 и т.п.).
// Делать это руками — мука; есть url.Values, который сам всё закодирует.
func getOpenMeteo() {
	base := "https://api.open-meteo.com/v1/forecast"

	params := url.Values{}
	params.Set("latitude", "55.75")
	params.Set("longitude", "37.61")
	params.Set("current", "temperature_2m,apparent_temperature")
	// temperature_2m — температура воздуха в градусах; apparent_temperature — ощущаемая температура
	params.Set("timezone", "Europe/Moscow")

	fullURL := base + "?" + params.Encode()
	fmt.Println("GET", fullURL)

	resp, err := apiClient.Get(fullURL)
	if err != nil {
		fmt.Println("ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("статус:", resp.Status)
		return
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("ответ:", string(body))
}

// --- 2. GET — стянуть один ресурс по id ---
//
// id уезжает прямо в путь URL, не в query-параметрах.
// Это разница не техническая, а смысловая: путь идентифицирует ресурс,
// query — задаёт дополнительные опции к запросу.
func getPost(id int) (*Post, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)

	resp, err := apiClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	// Decode читает прямо из потока — не нужно сначала ReadAll, а потом Unmarshal.
	var p Post
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}
	return &p, nil
}

// --- 3. POST — создать ресурс ---
//
// У POST/PUT тело запроса несёт данные. Для JSON это сериализованная структура
// плюс заголовок Content-Type: application/json — иначе сервер не поймёт формат.
// http.Client.Post принимает Content-Type вторым аргументом и тело — третьим.
func createPost(p Post) (*Post, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}

	resp, err := apiClient.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, fmt.Errorf("post: %w", err)
	}
	defer resp.Body.Close()

	// jsonplaceholder возвращает созданный объект с присвоенным id.
	var created Post
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}
	return &created, nil
}

// --- 4. PUT — заменить ресурс целиком ---
//
// Для всего, что не Get/Post, в стандартной библиотеке нет короткого хелпера.
// Делаем через http.NewRequest — это общий путь, через который под капотом
// идут и Get, и Post. Заголовки выставляем сами, через req.Header.Set.
func updatePost(p Post) (*Post, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", p.ID)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := apiClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}
	defer resp.Body.Close()

	var updated Post
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}
	return &updated, nil
}

// --- 5. DELETE — удалить ресурс ---
//
// Тела у DELETE-запроса обычно нет (в http.NewRequest передаём nil).
// Успех в REST принято обозначать кодом 200 или 204 (No Content).
func deletePost(id int) error {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}

	resp, err := apiClient.Do(req)
	if err != nil {
		return fmt.Errorf("do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	return nil
}

func apiDemo() {
	fmt.Println("=== GET с URL-параметрами (Open-Meteo) ===")
	getOpenMeteo()

	fmt.Println("\n=== GET одного ресурса (jsonplaceholder) ===")
	if p, err := getPost(1); err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Printf("получен: id=%d, title=%q\n", p.ID, p.Title)
	}

	fmt.Println("\n=== POST — создаём пост ===")
	newPost := Post{UserID: 1, Title: "Привет", Body: "из Go"}
	if p, err := createPost(newPost); err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Printf("создан: id=%d, title=%q\n", p.ID, p.Title)
	}

	fmt.Println("\n=== PUT — заменяем пост 1 ===")
	upd := Post{ID: 1, UserID: 1, Title: "Изменено", Body: "новый текст"}
	if p, err := updatePost(upd); err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Printf("обновлён: id=%d, title=%q\n", p.ID, p.Title)
	}

	fmt.Println("\n=== DELETE — удаляем пост 1 ===")
	if err := deletePost(1); err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("удалён пост 1")
	}
}
