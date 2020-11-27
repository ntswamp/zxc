package main

import (
	"fmt"
	"lgin/cmd/handler"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := "root:In198909@tcp(127.0.0.1:3306)/lgin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         logger.Default.LogMode(logger.Warn),
	})
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	/*
		row, _ := db.Raw("SELECT id,title FROM topic").Rows()
		for row.Next() {
			var id int
			var title string
			row.Scan(&id, &title)
			fmt.Println(id, title)
		}*/
	var tcs []handler.TopicClass
	db.Where(&handler.TopicClass{ClassName: "meetup"}).Find(&tcs)
	fmt.Println(tcs)
	defer sqlDB.Close()
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
