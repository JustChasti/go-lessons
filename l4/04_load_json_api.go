package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// --- структуры под ответ Open-Meteo (публичное API погоды, без ключа) ---
type CurrentWeather struct {
	Time        string  `json:"time"`
	Temperature float64 `json:"temperature_2m"`
	WindSpeed   float64 `json:"wind_speed_10m"`
}

type WeatherResponse struct {
	Latitude  float64        `json:"latitude"`
	Longitude float64        `json:"longitude"`
	Timezone  string         `json:"timezone"`
	Current   CurrentWeather `json:"current"`
}

// --- 1. Unmarshal из файла ---
// Читает product.json, который создал structsJSON(), и выводит структуру.
func productFromFile() {
	data, err := os.ReadFile("product.json")
	if err != nil {
		fmt.Println("не нашёл product.json — сначала запусти structsJSON()")
		return
	}

	var p Product
	if err := json.Unmarshal(data, &p); err != nil {
		fmt.Println("ошибка парсинга:", err)
		return
	}

	fmt.Println("--- прочитано из product.json ---")
	fmt.Printf("%+v\n", p)
}

// --- 2. HTTP-запрос + Unmarshal ---
// Публичное API погоды Open-Meteo. Ключ не нужен.
// Координаты 55.75/37.61 — Москва.
func fetchWeather() {
	url := "https://api.open-meteo.com/v1/forecast?latitude=55.75&longitude=37.61&current=temperature_2m,wind_speed_10m&timezone=Europe/Moscow"

	// таймаут на случай, если сервер долго не отвечает
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("плохой статус:", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ошибка чтения тела:", err)
		return
	}

	var w WeatherResponse
	if err := json.Unmarshal(body, &w); err != nil {
		fmt.Println("ошибка парсинга JSON:", err)
		return
	}

	fmt.Println("\n--- погода из Open-Meteo ---")
	fmt.Printf("координаты:   %.2f, %.2f\n", w.Latitude, w.Longitude)
	fmt.Printf("часовой пояс: %s\n", w.Timezone)
	fmt.Printf("время замера: %s\n", w.Current.Time)
	fmt.Printf("температура:  %.1f°C\n", w.Current.Temperature)
	fmt.Printf("ветер:        %.1f км/ч\n", w.Current.WindSpeed)
}

func loadJSONAPI() {
	productFromFile()
	fetchWeather()
}
