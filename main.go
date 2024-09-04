package main

import (
	"rest/db"
	"rest/routes"
	"github.com/gin-gonic/gin" //gin web framework(simplifies routing,handling HTTP requests,responding with JSON)
)

func main(){
	db.InitDB()
	server:=gin.Default() 

	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080
}

