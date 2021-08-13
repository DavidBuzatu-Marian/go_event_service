# Table of contents
- [Go Event Service](#go-event-service)
  * [Background](#background)
  * [Tech Stack](#tech-stack)

# Go Event Service
Event microservice for the [Datafy](https://github.com/DavidBuzatu-Marian/Datafy) project. It is concerned with the [Google Calendar API](https://developers.google.com/calendar) and adding events there.

## Background
This microservice is used to fetch the latest events added in the past hour and creating Google Calendar events for them using my personal account.

It uses Datafy's REST API to fetch the events, and it handles the creation and addition of events using the Google Calendar's API.

## Tech Stack
- [GoLang](https://golang.org)
- [Google Calendar API](https://developers.google.com/calendar)
