package main

import (
	"fmt"

	"github.com/DavidBuzatu-Marian/go_mongo"
)

func PrintEvents() {
	client := go_mongo.ConnectToMongo(config.MongoURI)
	events := go_mongo.CollectEvents(client)
	for _, val := range events {
		fmt.Println(val)
	}
}
