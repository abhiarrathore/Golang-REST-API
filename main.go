package main

import (
	"encoding/json"
	"fmt"
	"github.com/bb-tb-abhishek/gotutorialrest/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

var events = models.Events

func homeLink(w http.ResponseWriter, r * http.Request) {
	_, err := w.Write([]byte("Abhishek"))
	if err != nil {
		return
	}
}

func getAllEvents(w http.ResponseWriter, r * http.Request) {
	err := json.NewEncoder(w).Encode(events)
	if err != nil {
		return
	}
}

func getEvent(w http.ResponseWriter, r * http.Request) {
	eventID, _ := strconv.Atoi(mux.Vars(r)["id"])

	for _, event := range events {
		if event.ID == eventID {
			err := json.NewEncoder(w).Encode(event)
			if err != nil {
				return
			}
		}
	}
}

func createEvent(w http.ResponseWriter, r * http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	var newEvent models.Event

	if err != nil {}

	err2 := json.Unmarshal(reqBody, &newEvent)
	if err2 != nil {
		return 
	}

	events = append(events, newEvent)
}

func updateEvent(w http.ResponseWriter, r * http.Request) {
	eventID, _ := strconv.Atoi(mux.Vars(r)["id"])
	reqBody, err := ioutil.ReadAll(r.Body)
	var newEvent models.Event

	if err != nil {}

	err2 := json.Unmarshal(reqBody, &newEvent)
	if err2 != nil {
		return
	}

	for i, event := range events {
		if event.ID == eventID {
			event.Title = newEvent.Title
			event.Description = newEvent.Description
			events = append(events[:i], event)
			json.NewEncoder(w).Encode(event)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r * http.Request) {
	eventID, _ := strconv.Atoi(mux.Vars(r)["id"])

	for i, event := range events {
		if event.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "Event with ID %v is deleted", eventID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink).Methods("GET")
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", getEvent).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	fmt.Println("server listening on port 8000 yeah")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		return
	}
}
