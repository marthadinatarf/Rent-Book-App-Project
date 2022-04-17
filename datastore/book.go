package datastore

import (
	"fmt"
	"rent-book-app-project/schema"

	"gorm.io/gorm"
)

func InsertBuku(db *gorm.DB, newBook schema.Book) (schema.Book, error) {
	res := db.Create(&newBook)

	if res.Error != nil {
		fmt.Println("Insert Book Error : ", res.Error)
		return schema.Book{}, res.Error
	}

	return newBook, nil
}

func UpdateBuku(db *gorm.DB, updateBook schema.Book) (schema.Book, error) {
	res := db.Save(&updateBook)

	if res.Error != nil {
		fmt.Println("Update Book Error : ", res.Error)
		return schema.Book{}, res.Error
	}
	return updateBook, nil
}

func TampilkanBuku(db *gorm.DB) ([]schema.Book, error) {
	res := []schema.Book{}

	qry := db.Find(&res)

	if qry.Error != nil {
		fmt.Println("Error Select All User: ", qry.Error)
		return nil, qry.Error
	}

	return res, nil
}
