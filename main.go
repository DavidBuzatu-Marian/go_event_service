package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	MongoURI string
}

var config Config

func ReadConfig() {
	file, _ := os.Open("./config/default.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	config = Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error reading config file!")
		os.Exit(1)
	}
}

func main() {
	ReadConfig()
}
