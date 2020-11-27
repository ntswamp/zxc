package main

import (
	"lgin/cmd/handler"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

func main() {

	/*
		topic := handler.Topic{
			Title:      "shut up jenny!",
			ShortTitle: "mutejen",
			UserIp:     "127.0.0.222",
			Score:      3,
			Url:        "shutuup",
			Date:       time.Now(),
		}

		fmt.Println(db.Create(&topic).RowsAffected)
		fmt.Println(topic.Id)

			row, _ := db.Raw("SELECT id,title FROM topic").Rows()
			for row.Next() {
				var id int
				var title string
				row.Scan(&id, &title)
				fmt.Println(id, title)
			}

				var tcs []handler.TopicClass
				db.Where(&handler.TopicClass{ClassName: "meetup"}).Find(&tcs)
				fmt.Println(tcs)
	*/

	router := gin.Default()
	//validator stuff
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("topicurl", handler.TopicUrlValidator)
		v.RegisterValidation("topicssize", handler.TopicsSizeValidator)
	}

	v1 := router.Group("/topic")
	{
		//list all topic
		v1.GET("", handler.GetAllTopic())
		//get one topic
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
}
