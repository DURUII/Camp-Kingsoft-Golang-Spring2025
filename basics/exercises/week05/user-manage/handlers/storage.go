package handlers

import (
	"encoding/json"
	"os"
	"sync"
)

var (
	users []User
	mu    sync.RWMutex
)

const filePath = "user.json"

func initStorage() {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Create("user.json")
	}
	loadFromFile()
}

func loadFromFile() {
	data, _ := os.ReadFile(filePath)
	if len(data) == 0 {
		users = make([]User, 0)
		return
	}
	_ = json.Unmarshal(data, &users)
}

func saveToFile() error {
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func init() {
	initStorage()
}
