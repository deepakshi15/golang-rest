// event.go-logic that deals with storing event data in a database
package models

import (
	"rest/db"
	"time"
)

type Event struct {
	ID          int64
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

	stmt,err:=db.DB.Prepare(query) //prepares a sql statement,alternatively,also directly execute a statement via Exec()

	if err!=nil{
		return err
	}

	defer stmt.Close()

	//execute the prepared statement
	result,err:=stmt.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserID)
	if err!=nil{
		return err
	}

	id,err:=result.LastInsertId() //get the id of that event that was inserted
	e.ID=id
    return err
}

// function named GetAllEvents that returns a slice of 'Event' objects
func GetAllEvents() ([]Event,error) {
	query:= "SELECT *FRom events"
	rows, err:= db.DB.Query(query)
	if err!=nil{
		return nil,err
	}
	
	defer rows.Close()

	var events []Event

	for rows.Next(){
		var event Event
		err:=rows.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime) //pointer to id field in event struct

		if err!=nil{
			return nil, err
		}
		events=append(events, event)
	}

	return events,nil 
}
