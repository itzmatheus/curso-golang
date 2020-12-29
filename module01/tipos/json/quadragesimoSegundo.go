package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{
		ID:    1,
		Nome:  "Dell 15 7580",
		Preco: 3999.90,
		Tags:  []string{"dell", "notebook", "windows"},
	}

	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString := `
	{
		"id": 2,
		"nome": "Mousepad Gamer HyperX Fury LG",
		"preco": 179.90,
		"tags": [
			"hyperx",
			"mousepad",
			"gamer"
		]
	}
	`

	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2)

}
