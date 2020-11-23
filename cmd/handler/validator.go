package handler

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// TopicURL ...
func TopicURL(fl validator.FieldLevel) bool {
	if fl.StructFieldName() == "Top6ics" {
		fmt.Println(fl.StructFieldName())
		if matched, _ := regexp.MatchString(`^\w{4,10}$`, fl.Field().String()); matched {
			return true
		}
	}
	return false
}
