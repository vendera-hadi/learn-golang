package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// create a book
func Create(db *gorm.DB, Book *Book) (err error) {
	err = db.Create(Book).Error
	if err != nil {
		return err
	}
	return nil
}

//get books
func GetBooks(db *gorm.DB, Book *[]Book) (err error) {
	err = db.Find(Book).Error
	if err != nil {
		return err
	}
	return nil
}

//get book by id
func GetBook(db *gorm.DB, Book *Book, id string) (err error) {
	err = db.Where("id = ?", id).First(Book).Error
	if err != nil {
		return err
	}
	return nil
}

//update book
func Update(db *gorm.DB, Book *Book) (err error) {
	db.Save(Book)
	return nil
}

//delete book
func Delete(db *gorm.DB, Book *Book, id string) (err error) {
	db.Where("id = ?", id).Delete(Book)
	return nil
}
