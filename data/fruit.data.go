package data

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type Fruit struct {
	Name        string  `fake:"{fruit}"`
	Description string  `fake:"{loremipsumsentence:10}"`
	Price       float64 `fake:"{price:1,10}"`
}

func generateFruit() []string {
	var f Fruit
	gofakeit.Struct(&f)

	fruit := []string{}
	fruit = append(fruit, f.Name)
	fruit = append(fruit, f.Description)
	fruit = append(fruit, fmt.Sprintf("%.2f", f.Price))

	return fruit
}

func FruitList(length int) [][]string {
	var fruits [][]string

	for i := 0; i < length; i++ {
		f := generateFruit()
		fruits = append(fruits, f)
	}

	return fruits
}
