package main

import "fmt"

// Урок 6.5 — ДОМАШКА: «погодный сторож» в Telegram.
// Полное описание задачи, ссылки на документацию и подсказки — в README.md.
//
// Идея: один проход — стянули погоду, проверили пороги, если что-то превышено —
// отправили сообщение в Telegram

// === 1. Конфигурация ===
//
// Минимальный набор полей:
//   - City                 — для подписи в сообщении
//   - Latitude, Longitude  — Open-Meteo принимает координаты, не имя города
//   - WindThreshold        — порог по ветру (Open-Meteo по умолчанию даёт км/ч)
//   - RainThreshold        — порог по вероятности осадков, 0..100
//   - BotToken             — токен, который выдал @BotFather
//   - ChatID               — id чата, куда слать сообщение
//
// Структура может жить в этом файле, но интереснее — вынести её в подпакет
// "l6/config"

// === 2. fetchWeather — забрать данные с Open-Meteo ===
//
// Документация: https://open-meteo.com/en/docs
// Эндпоинт:     https://api.open-meteo.com/v1/forecast
// Параметры запроса:
//   latitude, longitude
//   current=wind_speed_10m,precipitation_probability
//   timezone=auto
//
// type WeatherNow struct {
// 	WindSpeed  float64
// 	RainChance float64
// }
//
// func fetchWeather(cfg Config) (WeatherNow, error) {
// 	return WeatherNow{}, nil
// }

// === 3. sendTelegram — отправить сообщение в чат ===
//
// Документация: https://core.telegram.org/bots/api#sendmessage
// Эндпоинт:     https://api.telegram.org/bot<TOKEN>/sendMessage
// Метод:        POST
//
// func sendTelegram(cfg Config, text string) error {
// 	return nil
// }

// === 4. shouldNotify — проверить пороги ===
//
// func shouldNotify(cfg Config, w WeatherNow) (bool, string) {
// 	return false, ""
// }

// === 5. practice — связать всё вместе ===

func practice() {
	fmt.Println("Домашка 6: реализуй задачи в 05_practice.go")

	// --- собранный сценарий ---
}
