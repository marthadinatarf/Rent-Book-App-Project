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

func MenuRegister(dbConn *gorm.DB) {
	var newUser schema.User
	fmt.Print("Masukkan Nama :")
	fmt.Scanln(&newUser.Nama)
	fmt.Print("Masukkan HP :")
	fmt.Scanln(&newUser.Hp)
	fmt.Print("Masukkan Alamat :")
	fmt.Scanln(&newUser.Alamat)
	fmt.Print("Masukkan Email :")
	fmt.Scanln(&newUser.Email)
	fmt.Print("Masukkan Password :")
	fmt.Scanln(&newUser.Password)

	res, err := datastore.InsertUser(dbConn, newUser)
	if err != nil {
		fmt.Println("terjadi sebuah kesalahan :", err)
	}
	fmt.Println(res.Nama, " Berhasil Didaftarkan")
}

func MenuLogin(dbConn *gorm.DB) {
	var userLogin schema.User
	var input int
	fmt.Print("Masukkan Email : ")
	fmt.Scanln(&userLogin.Email)
	fmt.Print("Masukkan Password : ")
	fmt.Scanln(&userLogin.Password)

	res, err := datastore.GetUserLogin(dbConn, userLogin.Email, userLogin.Password)
	if err != nil {
		fmt.Println("terjadi sebuah kesalahan :", err)
	}
	fmt.Println(res.Nama, "Berhasi Login")
	for input != 9 {
		fmt.Println("Halo, ", res.Nama)
		fmt.Println("1. Update Profil")
		fmt.Println("2. Pinjam Buku")
		fmt.Println("3. Update Buku")
		fmt.Println("4. Delete Buku")
		fmt.Println("5. Tampilkan Semua data buku")
		fmt.Println("9. Logout")
		fmt.Print("Masukkan pilihan menu :")
		fmt.Scanln(&input)
		if input == 1 {
			fmt.Print("Ganti Nama :")
			fmt.Scanln(&userLogin.Nama)

			userLogin, err = datastore.UpdateUser(dbConn, userLogin)
			if err != nil {
				fmt.Println("terjadi sebuah kesalahan :", err)
			}
		}
	}
}

func main() {

	// load .env as config database credentials
	config := ReadEnv()

	// make connection to mysql database
	dbConn := ConnectDB(config)

	//membuat database migration
	//db.AutoMigrate(&schema.User{}, &schema.Book{})

	// akses datastore
	// userAcc := datastore.UserDB{Db: db}
	// loginAcc := datastore.UserDB{Db: db}
	// updateAcc := datastore.UserDB{Db: db}

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
			MenuRegister(dbConn)
		case 2:
			MenuLogin(dbConn)
		case 3:
			fmt.Println("Delete Profile")
		case 99:
			fmt.Println("Terimakasih banyak sudah mencoba program kami.")
			return
		default:
			fmt.Println("Maaf, pilihan inputan menu tidak tersedia!")
		} //end switch
	} //end loop
} //end main
