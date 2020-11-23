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
		c.JSON(http.StatusOK, ConstructTopic(924, "Fuck with Loba"))
	}
}

// GetAllTopic ...
func GetAllTopic() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := QueryParam{}
		err := c.BindQuery(&query)
		if err != nil {
			c.String(400, "invalid parameters: %s", err.Error())
		} else {
			c.JSON(200, query)
		}
	}
}

// PostNewTopic ...
func PostNewTopic() gin.HandlerFunc {
	return func(c *gin.Context) {
		topic := Topic{}
		err := c.BindJSON(&topic)
		if err != nil {
			c.String(400, "invalid parameters: %s", err.Error())
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
			c.String(400, "invalid parameters: %s", err.Error())
		} else {
			c.JSON(200, topics)
		}
	}
}
