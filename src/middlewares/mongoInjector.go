package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spksoft/go-gin-mgo-restapi-starter-project/src/db"
	"github.com/spksoft/go-gin-mgo-restapi-starter-project/src/models/topicModel"
)

// InjectMongoDB is a function that will inject mongoclient to each model
func InjectMongoDB(c *gin.Context) {
	s := db.Session.Clone()
	c.Set("dbSession", s)
	database := s.DB(db.Mongo.Database)
	topicModel.SetMongoDB(database)
	c.Next()
}
