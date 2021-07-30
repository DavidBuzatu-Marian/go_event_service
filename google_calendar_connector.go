package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var tokenFile = "./config/token.json"

func getClient(config *oauth2.Config) *http.Client {
	token := getTokenAndSaveIfNotFoundLocally(config)
	return config.Client(context.Background(), token)
}

func getTokenAndSaveIfNotFoundLocally(config *oauth2.Config) *oauth2.Token {
	token, err := getTokenFromFile(tokenFile)
	if err != nil {
		token = getTokenFromWeb(config)
		saveToken(tokenFile, token)
	}
	return token
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}
	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return token
}

func getTokenFromFile(file string) (*oauth2.Token, error) {
	file_stream, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer file_stream.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(file_stream).Decode(token)
	return token, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	file_stream, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer file_stream.Close()
	json.NewEncoder(file_stream).Encode(token)
}

func TestCalendar() {
	ctx := context.Background()
	jsonKey := getJsonKeyFromConfigFile()
	config := getConfigFromJSON(jsonKey)
	client := getClient(config)

	calendarService := createCalendarService(client, ctx)
	addDummyEvent(calendarService)
}

func addDummyEvent(calendarService *calendar.Service) {
	event := createEvent("Some title", "2021-07-31T09:00:00-07:00", "2021-07-31T10:00:00-07:00")
	calendarId := "primary"
	event, err := calendarService.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
}

func createEvent(title string, startDate string, endDate string) *calendar.Event {
	return &calendar.Event{
		Summary: title,
		Start: &calendar.EventDateTime{
			DateTime: startDate,
		},
		End: &calendar.EventDateTime{
			DateTime: endDate,
		},
	}
}

func createCalendarService(client *http.Client, ctx context.Context) *calendar.Service {
	service, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	return service
}

func getConfigFromJSON(jsonKey []byte) *oauth2.Config {
	// If modifying these scopes, delete previously saved token.json.
	config, err := google.ConfigFromJSON(jsonKey, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return config
}

func getJsonKeyFromConfigFile() []byte {
	jsonKey, err := ioutil.ReadFile("./config/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	return jsonKey
}
