package main

import (
	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spksoft/go-gin-mgo-restapi-starter-project/src/db"
	"github.com/spksoft/go-gin-mgo-restapi-starter-project/src/middlewares"
)

func loadENV() {
	var err error
	if len(os.Args) == 2 {
		err = godotenv.Load(os.Args[1])
	} else {
		err = godotenv.Load()
	}
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func init() {
	loadENV()
	db.Connect()
}

func main() {
	r := gin.Default()
	r.Use(middlewares.InjectMongoDB)
	r.Use(middlewares.ErrorHandler)
	r = setRouters(r)
	endless.ListenAndServe(":"+os.Getenv("PORT"), r)
}
