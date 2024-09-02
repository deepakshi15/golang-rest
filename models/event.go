// event.go-logic that deals with storing event data in a database
package models

import (
	"rest/db"
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

// defines a method Save() on the Event type-takes no arguments other than the receiver e
func (e Event) Save() error {
	query:=`
	INSERT INTO events(name,description,location,dateTime,user_id) 
	VALUES(?,?,?,?,?)`
	stmt,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer stmt.Close()
	result,err:=stmt.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserID)
	if err!=nil{
		return err
	}
	id,err:=result.LastInsertId()
	e.ID=id
	events = append(events, e)
}

// function named GetAllEvents that returns a slice of 'Event' objects
func GetAllEvents() []Event {
	return events //returns the current contents of the events slice
}
