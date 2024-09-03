package main

import (
	"net/http" //deals with HTTP statuses and methods

	"rest/db"
	"rest/models"
	"github.com/gin-gonic/gin" //gin web framework(simplifies routing,handling HTTP requests,responding with JSON)
)

func main(){
	db.InitDB()
	server:=gin.Default() // Gin server with default middleware, which includes logging and recovery. gin.Default() sets up the server with sensible defaults.

	server.GET("/events",getEvents) //GET,POST,PUT,PATCH,DELETE
	//When a GET request is made to /events, the function getEvents is called to handle it.
	
	server.POST("/events",createEvent) //sets up a route to handle POST requests to the /events endpoint. When a client makes a POST request to '/events', the 'createEvent' function is called to process the request.

	server.Run(":8080") //localhost:8080
	//server will listen for incoming HTTP requests at http://localhost:8080.
}

//handler func for /events route
func getEvents(context *gin.Context){ //takes a context parameter of type '*gin.context',represents the context of the current HTTP request and response.
	events,err:=models.GetAllEvents() //import the model package and use here
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch events.Try again later"})
		return 
	}
	context.JSON(http.StatusOK, events) 
}

//function that handles creating a new event
func createEvent(context *gin.Context){ //takes a context parameter,which provides methods for working with the HTTP request and response
	var event models.Event //declares a variable event of type models.Event
	err:=context.ShouldBindJSON(&event)  //passes a pointer to the 'event' variable so that 'ShouldBindJSON' can fill it with parsed data
	//this method automatically parses incoming JSON payload from the request body and binds it to the 'event' struct.

	if err!=nil{ //checks if an error occured while binding the JSON data
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse request data"})
		return
	}//If there is an error, the server responds with an HTTP 400 Bad Request status code and a JSON message indicating that the request data could not be parsed.

	//After the JSON data is successfully parsed, the function manually sets the ID and UserID fields of the event object. 
	event.ID=1
	event.UserID=1

	err=event.Save()

	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not create event.Try again later"})
		return
	}

	//this sends a JSON response back to the client.
	context.JSON(http.StatusCreated,gin.H{"message":"Event created!","event":event})
	//HTTP 201 Created status code indicates that a resource (the event) was successfully created.
}