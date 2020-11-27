package handler

import (
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

// Topic ...
type Topic struct {
	Id         int       `json:"id" gorm:"PRIMARY_KEY"`
	Title      string    `json:"title" binding:"min=2,max=30"`
	ShortTitle string    `json:"stitle" binding:"nefield=Title"`
	UserIp     string    `json:"ip" binding:"ip"`
	Score      int       `json:"score" binding:"omitempty,lte=4"`
	Url        string    `json:"url" binding:"omitempty,topicurl"`
	Date       time.Time `json:"date" binding:"required"`
}

// TopicURLValidator ...
func TopicUrlValidator(fl validator.FieldLevel) bool {
	if _, ok := fl.Parent().Interface().(Topic); ok {
		if matched, _ := regexp.MatchString(`^\w{4,10}$`, fl.Field().String()); matched {
			return true
		}
	}
	return false
}

//Topics ...
type Topics struct {
	Topics []*Topic `json:"topics" binding:"gt=0,lte=3,dive"`
	Size   int      `json:"size" binding:"topicssize"`
}

//TopicsSizeValidator ...
func TopicsSizeValidator(fl validator.FieldLevel) bool {
	if topics, ok := fl.Parent().Interface().(*Topics); ok {
		if topics.Size == len(topics.Topics) {
			return true
		}
	}
	return false
}

// QueryParam ...
type QueryParam struct {
	Username string `json:"username" form:"username" binding:"required"`
	Page     int    `json:"page" form:"page" binding:"required"`
	Display  int    `json:"display" form:"display" binding:"required"`
}

// ConstructTopic ...
func ConstructTopic(id int, title string) Topic {
	return Topic{Id: id, Title: title}
}

// TopicClass...
type TopicClass struct {
	ClassId     int
	ClassName   string
	ClassRemark string
	ClassWater  string `gorm:"Column:classwater"`
}
