package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:In198909@tcp(127.0.0.1:3306)/lgin?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	_ = db
	/*
		router := gin.Default()
		//validator stuff
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterValidation("topicurl", handler.TopicURLValidator)
			v.RegisterValidation("topicssize", handler.TopicsSizeValidator)
		}

		v1 := router.Group("/topic")
		{
			//list all topic
			v1.GET("", handler.GetAllTopic())
			//show a specific topic
			v1.GET("/:topic_id", handler.GetOneTopic())

			//Login needed action
			v1.Use(handler.LoginCheck())
			{
				//create a new topic
				v1.POST("", handler.PostNewTopic())
				//post multiple topics
				router.POST("/topics", handler.PostNewTopics())
				//delete a topic
				v1.DELETE("/delete/:topic_id", handler.DeleteTopic())
			}
		}
		router.Run()
	*/
}
