package main

import (
	"encoding/json"
	"os"
	"time"
)

type SavedData struct {
	Horny map[string]time.Time
}

var savedData *SavedData

func loadData() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal(data, &savedData); err != nil {
		panic(err.Error())
	}
}
func saveData() {
	data, err := json.Marshal(savedData)
	if err != nil {
		panic(err.Error())
	}

	if err := os.WriteFile("data.json", data, os.ModePerm); err != nil {
		panic(err.Error())
	}
}
