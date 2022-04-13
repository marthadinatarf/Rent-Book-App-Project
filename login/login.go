package login

import (
	"fmt"
	"rent-book-app-project/entitas"

	"gorm.io/gorm"
)

type UserDB struct {
	Db *gorm.DB
}

func (u *UserDB) GetDataEmail(ambilEmail, ambilPassword string) ([]entitas.User, error) {
	res := []entitas.User{}

	if err := u.Db.Table("user").Where("email = ? and password = ?", ambilEmail, ambilPassword).Find(&res).Error; err != nil {
		fmt.Println("Terjadi kesalahan !!", err)
		return []entitas.User{}, err
	}
	return res, nil
}

// func LoginUser() {
// var emailGet, passGet string
// fmt.Println("Masukkan email : ")
// fmt.Scanf("%s", emailGet)
// fmt.Println("Masukkan Password : ")
// fmt.Scanf("%s", passGet)
// _, err := GetDataEmail(emailGet, passGet)
// GetDataEmail(emailGet, passwordGet)
// }
