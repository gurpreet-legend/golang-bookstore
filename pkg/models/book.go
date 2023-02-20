package models

import (
	"myGolangProjects/golang-bookstore/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	deleteBook, _ := GetBookById(Id)
	db.Delete(&deleteBook)
	return *deleteBook
}
