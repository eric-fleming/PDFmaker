package data

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type Car struct {
	Make  string  `fake:"{carmaker}"`
	Model string  `fake:"{carmodel}"`
	Type  string  `fake:"{cartype}"`
	Price float64 `fake:"{price:18000,70000}"`
}

func generateCar() []string {
	var c Car
	gofakeit.Struct(&c)

	car := []string{}
	car = append(car, c.Make)
	car = append(car, c.Model)
	car = append(car, c.Type)
	car = append(car, fmt.Sprintf("%.0f", c.Price))

	return car
}

func CarList(length int) [][]string {
	var cars [][]string

	for i := 0; i < length; i++ {
		c := generateCar()
		cars = append(cars, c)
	}

	return cars
}
