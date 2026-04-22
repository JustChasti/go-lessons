package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// --- структура с json-тегами ---
//
// json:"id"              — поле всегда в JSON под именем "id"
// json:"description,omitempty" — поле пропускается, если оно "пустое"
//
//	(пустая строка, 0, nil, пустой срез/мапа)
//
// *bool вместо bool      — чтобы различать "не задано" (nil) и false
type Product struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	InStock     bool     `json:"in_stock"`
}

func structsJSON() {
	// --- 1. минимально заполненный продукт ---
	// Description, Tags, InStock пустые — omitempty их уберёт из JSON
	p1 := Product{
		ID:    1,
		Name:  "Коврик для мыши",
		Price: 499.90,
	}

	data1, err := json.MarshalIndent(p1, "", "  ")
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Println("--- минимальный продукт (omitempty убрал пустые поля) ---")
	fmt.Println(string(data1))

	// --- 2. полностью заполненный продукт ---
	stock := false
	p2 := Product{
		ID:          2,
		Name:        "Клавиатура",
		Price:       3499,
		Description: "механическая, коричневые свитчи",
		Tags:        []string{"периферия", "геймерское"},
		InStock:     stock,
	}

	data2, err := json.MarshalIndent(p2, "", "  ")
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Println("\n--- полный продукт ---")
	fmt.Println(string(data2))

	// --- 4. сохранение в файл ---
	if err := os.WriteFile("product.json", data2, 0644); err != nil {
		fmt.Println("ошибка записи:", err)
		return
	}
	fmt.Println("\nсохранено в product.json")
}
