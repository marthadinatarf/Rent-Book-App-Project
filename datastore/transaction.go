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

func UpdateKembaliBuku(db *gorm.DB, updateTanggal string, id uint) (schema.Transactions, error) {
	res := schema.Transactions{}
	qry := db.Model(&res).Where("id = ?", id).Update("nama", updateTanggal)

	if qry.Error != nil {
		fmt.Println("Error Update User: ", qry.Error)
		return schema.Transactions{}, qry.Error
	}

	return res, nil
}
