package schema

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	Hp       string
	Alamat   string
	Email    string
	Password string
}

type Book struct {
	gorm.Model
	Judul       string
	Penerbit    string
	Penulis     string
	TahunTerbit int
}
