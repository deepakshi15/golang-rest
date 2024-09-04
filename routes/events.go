package routes

import (
	"context"
	"net/http" //deals with HTTP statuses and methods
	"rest/models"
	"strconv"
	"github.com/gin-gonic/gin" //gin web framework(simplifies routing,handling HTTP requests,responding with JSON)
)

func getEvents(context *gin.Context){ 
	events,err:=models.GetAllEvents() 
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch events.Try again later"})
		return 
	}
	context.JSON(http.StatusOK, events) 
}

//function for getting a single event now-handler function for the GET request
func getEvent(context *gin.Context){ 
	eventID,err:=strconv.ParseInt(context.Param("id"),10,64) 
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse event id."}) 
		return
	}
	event,err:=models.GetEventByID(eventID) 

	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch event."}) 
		return
	}
	context.JSON(http.StatusOK,event) 
}

func createEvent(context *gin.Context){ 
	var event models.Event 
	err:=context.ShouldBindJSON(&event)  

	if err!=nil{ 
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse request data"})
		return
	}

	event.ID=1
	event.UserID=1

	err=event.Save()

	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not create event.Try again later"})
		return
	}

	context.JSON(http.StatusCreated,gin.H{"message":"Event created!","event":event})
}

func updateEvent(context *gin.Context){
	eventID,err:=strconv.ParseInt(context.Param("id"),10,64)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse event"})
		return
	}

	_,err=models.GetEventByID(eventID)

	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":""})
	}

	var updateEvent models.Event
	err=context.ShouldBindJSON(&updateEvent)

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse event data"})
		return 
	}
	updateEvent.ID=eventID
	err=updateEvent.update()

	if err!=nil{
		context.JSON((http.StatusInternalServerError,gin.H{"message":"Could not update event."}))
		return
	}	
	context.JSON(http.StatusOk,gin.H{"message":"Event updated successfully"})
}