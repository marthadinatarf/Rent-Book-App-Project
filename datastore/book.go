package datastore

import (
	"fmt"
	"rent-book-app-project/schema"

	"gorm.io/gorm"
)

func InsertBukur(db *gorm.DB, newBook schema.Book) (schema.Book, error) {
	res := db.Create(&newBook)

	if res.Error != nil {
		fmt.Println("Insert Book Error : ", res.Error)
		return schema.Book{}, res.Error
	}

	return newBook, nil
}
