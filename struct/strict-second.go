package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Persone struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"Дата рождения"`
}

func main() {
	client := Persone{Name: "Test", Email: "test@test.my"}

	jsClient, err := json.Marshal(client)

	if err != nil {
		log.Fatal("unable marshal to json")
	}

	fmt.Println(string(jsClient))

	req := struct {
		NameContains string `json:"name_contains"`
		Offset       int    `json:"offset"`
		Limit        int    `json:"limit"`
	}{
		NameContains: "Иван",
		Limit:        50,
	}

	reqRaw, _ := json.Marshal(req)
	fmt.Println(string(reqRaw))
}
