package handler

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Topic ...
type Topic struct {
	ID         int    `json:"id"`
	Title      string `json:"title" binding:"min=2,max=30"`
	ShortTitle string `json:"stitle" binding:"nefield=Title"`
	UserIP     string `json:"ip" binding:"ip"`
	Score      int    `json:"score" binding:"omitempty,lte=4"`
	URL        string `json:"url" binding:"omitempty,topicurl"`
}

// TopicURLValidator ...
func TopicURLValidator(fl validator.FieldLevel) bool {
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
	return Topic{ID: id, Title: title}
}

// TopicClass...
type TopicClass struct {
	ClassId     int
	ClassName   string
	ClassRemark string
	ClassWater  string `gorm:"Column:classwater"`
}
