package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gin-mgo-restapi-starter-project/src/controllers/topicController"
)

// Routers is handler
func setRouters(r *gin.Engine) *gin.Engine {
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"body": "ok",
		})
	})

	v1 := r.Group("/api/v1")
	{
		topics := v1.Group("/topics")
		{
			topics.GET("/:_id", topicController.GetTopicByID)
			topics.GET("/", topicController.GetTopics)
			topics.POST("/", topicController.CreateTopic)
			topics.DELETE("/:_id", topicController.DeleteTopic)
		}
	}
	return r
}
