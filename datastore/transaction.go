package datastore

import (
	"fmt"
	"rent-book-app-project/schema"

	"gorm.io/gorm"
)

func InsertPinjamBUKU(db *gorm.DB, newPinjam schema.Transactions) (schema.Transactions, error) {
	res := db.Create(&newPinjam)

	if res.Error != nil {
		fmt.Println("Insert Book Error : ", res.Error)
		return schema.Transactions{}, res.Error
	}

	return newPinjam, nil
}

func InsertKembaliBuku(db *gorm.DB, newKembali schema.Transactions) (schema.Transactions, error) {
	res := db.Create(&newKembali)

	if res.Error != nil {
		fmt.Println("Insert Book Error : ", res.Error)
		return schema.Transactions{}, res.Error
	}

	return newKembali, nil
}
