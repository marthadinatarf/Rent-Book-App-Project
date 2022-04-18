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
	qry := db.Model(&res).Where("id = ?", id).Update("tanggal_kembali", updateTanggal)

	if qry.Error != nil {
		fmt.Println("Error Update User: ", qry.Error)
		return schema.Transactions{}, qry.Error
	}

	return res, nil
}
func GetTransactionById(db *gorm.DB, id uint) (schema.Transactions, error) {
	trans := schema.Transactions{}
	if err := db.Select("tanggal_pinjam", "user_id", "book_id").Where("id = ?", id).Find(&trans).Error; err != nil {
		fmt.Println("Get Book By Id Error:", err)
		return trans, err
	}
	return trans, nil
}
