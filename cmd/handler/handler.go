package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginCheck ...
func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, Loggedin := c.GetQuery("token"); Loggedin {
			c.Next()
		} else {
			c.String(http.StatusUnauthorized, "login token undetected.")
			c.Abort()
		}
	}
}

// GetOneTopic ...
func GetOneTopic() gin.HandlerFunc {
	return func(c *gin.Context) {

		tid := c.Param("topic_id") // capture params specified in source code i.e main.go /:topic_id
		topic := Topic{}
		DB.Find(&topic, tid)
		c.JSON(200, topic)
		//} else {
		//	c.String(http.StatusNotFound, "topic not found")
	}
}

// GetAllTopic ...
func GetAllTopic() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, ConstructTopic(924, "All topics are listed. Let's F*ck with Loba now"))
	}
}

// PostNewTopic ...
func PostNewTopic() gin.HandlerFunc {
	return func(c *gin.Context) {
		topic := Topic{}
		err := c.BindJSON(&topic)
		if err != nil {
			c.String(400, "invalid parameters:\n%s", err.Error())
		} else {
			c.JSON(200, topic)
		}
	}
}

// DeleteTopic ...
func DeleteTopic() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "topic was deleted.")
	}
}

// PostNewTopics ...
func PostNewTopics() gin.HandlerFunc {
	return func(c *gin.Context) {
		topics := Topics{}
		err := c.BindJSON(&topics)
		if err != nil {
			c.String(400, "invalid parameters:\n%s", err.Error())
		} else {
			c.JSON(200, topics)
		}
	}
}
