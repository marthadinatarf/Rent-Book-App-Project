package datastore

import (
	"fmt"
	"rent-book-app-project/schema"

	"gorm.io/gorm"
)

func InsertUser(db *gorm.DB, newUser schema.User) (schema.User, error) {
	res := db.Create(&newUser)

	if res.Error != nil {
		fmt.Println("Insert User Error : ", res.Error)
		return schema.User{}, res.Error
	}

	return newUser, nil
}

func GetAllUser(db *gorm.DB) ([]schema.User, error) {
	res := []schema.User{}

	qry := db.Find(&res)

	if qry.Error != nil {
		fmt.Println("Error Select All User: ", qry.Error)
		return nil, qry.Error
	}

	return res, nil
}

func GetUserLogin(db *gorm.DB, email, password string) (schema.User, error) {
	usr := schema.User{}

	qry := db.Select("id", "nama", "email", "password").Where("email = ? and password = ?", email, password).Find(&usr)

	if qry.Error != nil {
		fmt.Println("Error Select User Login: ", qry.Error)
		return schema.User{}, qry.Error
	}

	return usr, nil
}

func UpdateUser(db *gorm.DB, updateNama string, id uint) (schema.User, error) {
	user := schema.User{}
	qry := db.Model(&user).Where("id = ?", id).Update("nama", updateNama)

	if qry.Error != nil {
		fmt.Println("Error Update User: ", qry.Error)
		return schema.User{}, qry.Error
	}

	return user, nil
}

func DeleteUser(db *gorm.DB, deleteUser schema.User) (schema.User, error) {
	res := db.Delete(&deleteUser)

	if res.Error != nil {
		fmt.Println("Error delete user : ", res.Error)
		return schema.User{}, res.Error
	}
	return deleteUser, nil
}
