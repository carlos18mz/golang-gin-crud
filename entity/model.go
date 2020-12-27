package entity

type Person struct {
	ID        uint64
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}

type Video struct {
	ID          uint64 `json:"id" gorm:"primary_key; auto_increment"`
	Title       string `json:"title" binding:"max=80" validate:"is-cool"`
	Description string `json:"description" binding:"max=100"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
