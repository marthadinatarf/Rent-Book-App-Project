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

func GetBukuById(db *gorm.DB, id uint) (schema.Book, error) {
	book := schema.Book{}
	if err := db.Select("judul", "penerbit", "penulis", "tahun_terbit").Where("id = ?", id).Find(&book).Error; err != nil {
		fmt.Println("Get Book By Id Error:", err)
		return book, err
	}
	return book, nil
}

func UpdateBuku(db *gorm.DB, id uint, books schema.Book) (schema.Book, error) {
	query := db.Model(&books).Where("id = ?", id).Updates(schema.Book{
		Judul:       books.Judul,
		Penerbit:    books.Penerbit,
		Penulis:     books.Penulis,
		TahunTerbit: books.TahunTerbit,
	})
	if query.Error != nil {
		fmt.Println("Update Book Error : ", query.Error)
		return books, query.Error
	}
	return books, nil
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
