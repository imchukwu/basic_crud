package models

import (
	"github.com/imchukwu/book_management_system/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Publication string `json:"publication"`
	Author      string `json:"author"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) *Book {
	var book *Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
