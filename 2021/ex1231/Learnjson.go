package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	type Fruit struct {
		Name     string `json:"name"`
		PriceTag string `json:"priceTag"`
	}
	//解析具有动态Key的对象
	type FruitBasket struct {
		Name    string           `json:"name"`
		Fruit   map[string]Fruit `json:"fruit"`
		Id      int64            `json:"id"`
		Created time.Time        `json:"created"`
	}
	jsonData := []byte(`
    {
        "Name": "Standard",
        "Fruit" : {
              "1": {
                    "name": "Apple",
                    "priceTag": "$1"
              },
              "2": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              },
              "3": {
                    "name": "Watermelon",
                    "priceTag": "$10.5"
              },
              "4": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              },
              "5": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              },
              "6": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              },
              "7": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              },
              "8": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              },
              "9": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              },
               "10": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              },
               "11": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              },
               "12": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              },            
	           "13": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              },
               "14": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              }

        },
        "id": 999,
        "created": "2018-04-09T23:00:00Z"
    }`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range basket.Fruit {
		fmt.Println(item.Name, item.PriceTag, "basket3:", basket.Fruit)
	}
}
