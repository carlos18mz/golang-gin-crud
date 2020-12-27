package repository

import (
	"fmt"

	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

type database struct {
	connection gorm.DB
}

func NewVideoRepository() VideoRepository {

	user := "uxzs73sahif4nyn8"
	pass := "FeNLsTS7r6QZDwodnLyy"
	host := "bocf53vcqoztozmdfvlh-mysql.services.clever-cloud.com"
	dbname := "bocf53vcqoztozmdfvlh"
	port := 3306

	dsn := fmt.Sprintf("mysql://%d:%d@%d:%d/%d?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&entity.Video, &entity.Person)
	return &database{
		connection: db,
	}
}
