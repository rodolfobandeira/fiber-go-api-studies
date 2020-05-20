package main

import (
	"github.com/gofiber/fiber"
)

type Pagination struct {
	CurrentPage uint16
	TotalPages  uint16
	NextPage    uint16
	PrevPage    uint16
}

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

type ResponseData struct {
	Pagination Pagination
	Cars       []Car
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

	pagination1 := Pagination{
		CurrentPage: 1,
		TotalPages:  1,
		NextPage:    0,
		PrevPage:    0,
	}

	app.Get("/cars", func(c *fiber.Ctx) {
		c.JSON(
			ResponseData{
				pagination1,
				[]Car{car1, car2},
			},
		)
	})

	app.Listen(3000)
}

/*
{
    "Pagination": {
        "CurrentPage": 1,
        "TotalPages": 1,
        "NextPage": 0,
        "PrevPage": 0
    },
    "Cars": [
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
}
*/
