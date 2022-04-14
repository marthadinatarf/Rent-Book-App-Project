package datastore

import (
	"fmt"
	"rent-book-app-project/schema"

	"gorm.io/gorm"
)

type UserDB struct {
	Db *gorm.DB
}

func (u *UserDB) RegisterUser(addUser schema.User) (schema.User, error) {
	if err := u.Db.Select("Nama", "Hp", "Alamat", "Email", "Password").Create(&addUser).Error; err != nil {
		fmt.Println("Terjadi kesalahan saat register user", err)
		return addUser, err
	}
	return addUser, nil
}
