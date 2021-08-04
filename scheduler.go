package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Event []struct {
	ID         string    `json:"_id"`
	Name       string    `json:"name"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	Location   string    `json:"location"`
	Directions string    `json:"directions"`
	Details    string    `json:"details"`
	DateAdded  time.Time `json:"date_added"`
	CalendarID string    `json:"calendar_id"`
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func Schedule(repeatInterval time.Duration) {
	for {
		events := new(Event)
		err := getPastHourEvents("http://localhost:8080/api/event/info/latest", events)
		if err != nil {
			log.Fatal(err)
		}
		for _, val := range *events {
			fmt.Println(val)
		}
		AddEvents(events)
		<-time.After(repeatInterval * time.Second)
	}

}

func getPastHourEvents(url string, target interface{}) error {
	response, err := httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(&target)
}
