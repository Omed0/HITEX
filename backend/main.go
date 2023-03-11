package main

import (
	//"fmt"

	"log"
	"emad.com/config"
	"emad.com/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init()  {

	 err := godotenv.Load()

     if err != nil {
     log.Fatal("Error loading .env file")
}
	
}

func main() {

    config.ConnectDB()

	router := gin.Default()

	routes.SetUpRowters(router)

	router.Run() 
	
}