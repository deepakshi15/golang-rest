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
var events = []Event{} 

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
    return err
}

func GetAllEvents() ([]Event,error) {
	query:= "SELECT *FROM events"
	rows, err:= db.DB.Query(query)
	if err!=nil{
		return nil,err
	}
	
	defer rows.Close()

	var events []Event

	for rows.Next(){
		var event Event
		err:=rows.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID) 

		if err!=nil{
			return nil, err
		}
		events=append(events, event)
	}

	return events,nil 
}

//The function takes an id of type int64, representing the unique identifier for an event
func GetEventByID(id int64) (*Event,error){ //returns pointer to an event struct and an error
	query:="SELECT *FROM events WHERE id=?" //defines the SQL query to select all columns (*) from the events table where the id column matches the given id parameter,? is a placeholder for parameterized queries, which prevents SQL injection.
	row:=db.DB.QueryRow(query,id) //query and retrieves a single row from the events table where the event id matches the provided id.

	var event Event //variable event of type Event is created to hold the fetched event data.
	//Scan method maps the columns returned by the query into the fields of the event struct. 
	err:=row.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID) //&event.ID will store the value from the id column in the ID field of the event struct.
	if err!=nil{
		return nil, err //If the Scan method fails, the function returns nil for the event and the error that occurred.
	}
	return &event,nil //If the query successfully retrieves the event, the function returns a pointer to the event and nil for the error.
}

func (event Event) update() error{
	query:=`
	UPDATE events
	SET name=?, description=?,location=?,dateTime=?
	WHERE id=?
	`
	stmt,err:=db.DB.Prepare(query)

	if err!=nil{
		return err
	}

	defer stmt.Close()

	_,err=stmt.Exec(event.Name,event.Description,event.Location,event.DateTime,event.UserID)
	return err
}
