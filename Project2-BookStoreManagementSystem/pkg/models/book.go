package models

import (
	"log"

	"github.com/ankit/Project2-BookManagementSystem/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	/*The gorm.Model struct provides common fields
	like ID, CreatedAt, UpdatedAt, and DeletedAt.*/
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Author      string `gorm:"column:author" json:"author"`
	Publication string `gorm:"column:publication" json:"publication"`
	// ...
}

func init() {
	config.Connect()
	db = config.GetDB()
	//db.AutoMigrate(&Book{})
	err := db.AutoMigrate(&Book{}).Error
	if err != nil {
		log.Fatal("Error during database migration:", err)
	}
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
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

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Find(book)
	return book
}
