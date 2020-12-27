package entity

import "time"

type Person struct {
	ID        uint64 `json:"id" gorm:"primary_key;auto_increment"`
	FirstName string `json:"firstname" binding:"max=100" gorm:"type:varchar(100)"`
	LastName  string `json:"lastname" binding:"max=100" gorm:"type:varchar(100)"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email" gorm:"type:varchar(256)"`
}

type Video struct {
	ID          uint64    `json:"id" gorm:"primary_key; auto_increment"`
	Title       string    `json:"title" binding:"max=100" validate:"is-cool" gorm:"type:varchar(100)"`
	Description string    `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:varchar(256); UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"type:time" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"type:time" json:"updated_at"`
}
