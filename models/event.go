// event.go-logic that deals with storing event data in a database
package models

import (
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

// declares a global variable named 'events'
var events = []Event{} //empty slice of 'event' type objects.slice is initialized but does not contain any elements initially.

// save method i.e attached to this event struct that will save such an event to the database
func (e Event) Save() {
	events = append(events, e)
}

// function named GetAllEvents that returns a slice of 'Event' objects
func GetAllEvents() []Event {
	return events //returns the current contents of the events slice
}
