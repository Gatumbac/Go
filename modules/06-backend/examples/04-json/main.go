package main

import (
	"encoding/json"
	"fmt"
)

type City struct {
	Name       string `json:"name"`
	Population int    `json:"population"`
}

func main() {
	data, err := json.Marshal(City{Name: "Loja", Population: 200})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
