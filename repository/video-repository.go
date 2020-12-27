package repository

import (
	"crud-gin/entity"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {

	user := "uxzs73sahif4nyn8"
	pass := "FeNLsTS7r6QZDwodnLyy"
	host := "bocf53vcqoztozmdfvlh-mysql.services.clever-cloud.com"
	dbname := "bocf53vcqoztozmdfvlh"
	port := "3306"

	//dsn := "mysql://uxzs73sahif4nyn8:FeNLsTS7r6QZDwodnLyy@bocf53vcqoztozmdfvlh-mysql.services.clever-cloud.com:3306/bocf53vcqoztozmdfvlh?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	//return db.Migrator().CurrentDatabase()
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	sqlDB, _ := db.connection.DB()

	err := sqlDB.Close()

	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
