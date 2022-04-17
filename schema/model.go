package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	Hp       string
	Alamat   string
	Email    string
	Password string
	Books    []Book
}

type Book struct {
	gorm.Model
	Judul        string
	Penerbit     string
	Penulis      string
	TahunTerbit  int
	UserID       uint
	Transactions Transactions
}

type Transactions struct {
	gorm.Model
	UserID      uint
	BookID      uint
	tgl_pinjam  time.Time
	tgl_kembali time.Time
}
