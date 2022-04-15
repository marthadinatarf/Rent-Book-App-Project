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
	// db.AutoMigrate(&schema.User{}, &schema.Book{})

	// akses datastore
	userAcc := datastore.UserDB{Db: db}
	// bookAcc := datastore.BookDB{Db: db}

	var pilihan int
	for pilihan != 99 {
		fmt.Println("")
		fmt.Println("======= Menu Utama =======")
		fmt.Println("1.\tRegister")
		fmt.Println("2.\tLogin")
		fmt.Println("3.\tLihat Daftar Buku")
		fmt.Println("99.\tKeluar")
		fmt.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+")
		fmt.Print("input menu pilihan : ")
		fmt.Scanln(&pilihan)
		fmt.Println("")
		switch pilihan {
		case 1:
			var users schema.User
			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&users.Nama)
			fmt.Print("Masukkan nomor hp: ")
			fmt.Scanln(&users.Hp)
			fmt.Print("Masukkan alamat: ")
			fmt.Scanln(&users.Alamat)
			fmt.Print("Masukkan email: ")
			fmt.Scanln(&users.Email)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&users.Password)

			res, err := userAcc.RegisterUser(
				schema.User{
					Nama:     users.Nama,
					Hp:       users.Hp,
					Alamat:   users.Alamat,
					Email:    users.Email,
					Password: users.Password,
				})
			if err != nil {
				fmt.Println(err)

			}
			fmt.Println(res)
		case 2:
			fmt.Println("Pilihan menu ke 2")
		case 3:
			fmt.Println("Pilihan menu ke 3")
		case 99:
			fmt.Println("Terimakasih banyak sudah mencoba program kami.")
			return
		default:
			fmt.Println("Maaf, pilihan inputan menu tidak tersedia!")
		} //end switch
	} //end loop
} //end main
