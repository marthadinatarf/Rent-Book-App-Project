package datastore

import (
	"fmt"
	"rent-book-app-project/schema"

	"gorm.io/gorm"
)

type UserDB struct {
	Db *gorm.DB
}

func (u *UserDB) GetDataEmail(ambilEmail, ambilPassword string) ([]schema.User, error) {
	res := []schema.User{}

	if err := u.Db.Table("user").Where("email = ? and password = ?", ambilEmail, ambilPassword).Find(&res).Error; err != nil {
		fmt.Println("Terjadi kesalahan !!", err)
		return []schema.User{}, err
	}
	return res, nil
}

func (u *UserDB) RegisterUser(addUser schema.User) (schema.User, error) {
	if err := u.Db.Create(&addUser).Error; err != nil {
		fmt.Println("Terjadi kesalahan saat register user", err)
		return addUser, err
	}
	return addUser, nil
}

func LoginUser() {
	var emailGet, passGet string
	fmt.Println("Masukkan email : ")
	fmt.Scan(&emailGet)
	fmt.Println("Masukkan Password : ")
	fmt.Scan(&passGet)
	if emailGet == "admin" && passGet == "admin" {
		fmt.Println("berhasil login")
	} else {
		fmt.Println("Email dan password salah")
	}
}
