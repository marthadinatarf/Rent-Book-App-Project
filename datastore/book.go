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

func UpdateBuku(db *gorm.DB, updateJudul schema.Book) (schema.Book, error) {
	res := db.Save(&updateJudul)

	if res.Error != nil {
		fmt.Println("Update Book Error : ", res.Error)
		return schema.Book{}, res.Error
	}
	return updateJudul, nil
}

func TampilkanBuku(db *gorm.DB) ([]schema.Book, error) {
	res := []schema.Book{}

	qry := db.Find(&res)

	if qry.Error != nil {
		fmt.Println("Terjadi kesalahan saat menampilkan buku: ", qry.Error)
		return nil, qry.Error
	}

	return res, nil
}

func DeleteBuku(db *gorm.DB, deleteBook schema.Book) (schema.Book, error) {
	res := db.Delete(&deleteBook)

	if res.Error != nil {
		fmt.Println("Terjadi kesalahan saat delete buku : ", res.Error)
		return schema.Book{}, res.Error
	}
	return deleteBook, nil
}
