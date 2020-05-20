# fiber-go-api-studies
Studying Filber GO as REST API through some examples (gofiber.io)

`go run main.go`

`http://localhost:3000/cars`

```json
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
```
