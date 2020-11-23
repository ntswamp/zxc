package handler

// Topic ...
type Topic struct {
	ID         int    `json:"id"`
	Title      string `json:"title" binding:"min=2,max=30"`
	ShortTitle string `json:"stitle" binding:"nefield=Title"`
	UserIP     string `json:"ip" binding:"ip"`
	Score      int    `json:"score" binding:"omitempty,lte=4"`
	URL        string `json:"url" binding:"omitempty" validate:"topicurl"`
}

//Topics ...
type Topics struct {
	Topics []Topic `json:"topics" binding:"gt=0,lte=3,dive"`
	Size   int     `json:"size"`
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
