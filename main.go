package main

import (
	"github.com/gofiber/fiber"
)

type Car struct {
	Name   string
	Year   uint16
	Engine Engine
}

type Engine struct {
	Cilinders uint8
	Power     float32
	GasType   string
}

func main() {
	app := fiber.New()

	car1 := Car{
		Name: "Jeep",
		Year: 1951,
		Engine: Engine{
			Cilinders: 6,
			Power:     3.0,
			GasType:   "Diesel",
		},
	}

	car2 := Car{
		Name: "Dodge",
		Year: 1975,
		Engine: Engine{
			Cilinders: 8,
			Power:     3.2,
			GasType:   "Gas",
		},
	}

	var Cars = []Car{car1, car2}

	app.Get("/car", func(c *fiber.Ctx) {
		data := Cars
		c.JSON(data)
	})

	app.Listen(3000)
}

/*
[
    {
        "Name": "Jeep",
        "Year": 1951,
        "Engine": {
            "Cilinders": 6,
            "Power": 3,
            "GasType": "Diesel"
        }
    },
    {
        "Name": "Dodge",
        "Year": 1975,
        "Engine": {
            "Cilinders": 8,
            "Power": 3.2,
            "GasType": "Gas"
        }
    }
]
*/
