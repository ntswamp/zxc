package main

import (
	"lgin/cmd/handler"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()
	//validator stuff
	validator := validator.New()
	validator.RegisterValidation("topicurl", handler.TopicURL)

	v1 := router.Group("/v1/topic")
	{
		//list all topic
		v1.GET("", handler.GetAllTopic())
		//list a specific topic
		v1.GET("/:topic_id", handler.GetOneTopic())

		//Login needed action
		v1.Use(handler.LoginCheck())
		{
			//create a new topic
			v1.POST("", handler.PostNewTopic())
			//post multiple topics
			router.POST("/v1/topics", handler.PostNewTopics())
			//delete a topic
			v1.DELETE("/deletetopic/:topic_id", handler.DeleteTopic())
		}
	}
	router.Run()
}
