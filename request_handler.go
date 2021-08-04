package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/api/calendar/v3"
)

type CalendarEventID struct {
	CalendarId string `json:"calendar_id"`
}

func GetPastHourEvents(url string, target interface{}) error {
	response, err := httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(&target)
}

func UpdateEvent(url string, event *calendar.Event) error {
	json := convertToJson(event)
	request := createRequest(url, json)
	_, err := httpClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func createRequest(url string, json []byte) *http.Request {
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(json))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	return request
}

func convertToJson(event *calendar.Event) []byte {
	payload := CalendarEventID{event.Id}
	json, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	return json
}
