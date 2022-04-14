package main

import (
	"fmt"
	"os"
	"rent-book-app-project/datastore"
	"rent-book-app-project/schema"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int16
	DB       string
}

func ReadEnv() Config {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("ERROR LOAD FILE", err)
	}
	res := Config{}
	res.Username = os.Getenv("USER_NAME")
	res.DB = os.Getenv("DATABASE")
	res.Password = os.Getenv("PASSWORD")
	res.Host = os.Getenv("HOST")
	intCon, _ := strconv.Atoi(os.Getenv("PORT"))
	res.Port = int16(intCon)
	return res
}

func ConnectDB(cfg Config) *gorm.DB {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB)

	db, err := gorm.Open(mysql.Open(dbConn), &gorm.Config{})
	if err != nil {
		fmt.Println("Terjadi kesalahan saat koneksi database", err)
		return nil
	}
	return db
}

func main() {

	// load .env as config database credentials
	config := ReadEnv()

	// make connection to mysql database
	db := ConnectDB(config)

	//membuat database migration
	//db.AutoMigrate(&schema.User{}, &schema.Book{})

	// akses user datastore
	userAcc := datastore.UserDB{Db: db}

	var pilihan int
	for pilihan != 4 {
		fmt.Println("====== Menu Utama ======")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Lihat Daftar Buku")
		fmt.Println("4. Keluar")
		fmt.Println("Masukkan pilihan : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			var users schema.User
			fmt.Scan(&users.Nama)
			fmt.Scan(&users.Hp)
			fmt.Scan(&users.Alamat)
			fmt.Scan(&users.Email)

			res, err := userAcc.RegisterUser(schema.User{Nama: users.Nama, Hp: users.Hp, Alamat: users.Alamat, Email: users.Email, Password: users.Password})
			if err != nil {
				fmt.Println(err)

			}
			fmt.Println(res)
		}
	}

	// switch pilihan {
	// case 1:
	// 	var users schema.User
	// 	fmt.Scan(&users.Nama)
	// 	fmt.Scan(&users.Hp)
	// 	fmt.Scan(&users.Alamat)
	// 	fmt.Scan(&users.Email)

	// 	res, err := userAcc.RegisterUser(schema.User{Nama: users.Nama, Hp: users.Hp, Alamat: users.Alamat, Email: users.Email, Password: users.Password})
	// 	if err != nil {
	// 		fmt.Println(err)

	// 	}
	// 	fmt.Println(res)
	// }

	// var emailGet, passGet string
	// fmt.Println("Masukkan email : ")
	// fmt.Scan(&emailGet)
	// fmt.Println("Masukkan Password : ")
	// fmt.Scan(&passGet)
	// if emailGet == "admin" && passGet == "admin" {
	// 	fmt.Println("berhasil login")
	// } else {
	// 	fmt.Println("Email dan password salah")
	// }

	//GetDataEmail(emailGet, passwordGet)

}
