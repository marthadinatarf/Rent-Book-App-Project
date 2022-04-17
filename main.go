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
	fmt.Print("Masukkan Nama : ")
	fmt.Scanln(&newUser.Nama)
	fmt.Print("Masukkan HP : ")
	fmt.Scanln(&newUser.Hp)
	fmt.Print("Masukkan Alamat : ")
	fmt.Scanln(&newUser.Alamat)
	fmt.Print("Masukkan Email : ")
	fmt.Scanln(&newUser.Email)
	fmt.Print("Masukkan Password : ")
	fmt.Scanln(&newUser.Password)

	res, err := datastore.InsertUser(dbConn, newUser)
	if err != nil {
		fmt.Println("terjadi sebuah kesalahan :", err)
	}
	fmt.Println(res.Nama, "Berhasil Didaftarkan")
}

func MenuLogin(dbConn *gorm.DB) {
	// var userLogin schema.User
	var email, password string
	var input int
	fmt.Print("Masukkan Email : ")
	fmt.Scanln(&email)
	fmt.Print("Masukkan Password : ")
	fmt.Scanln(&password)

	res, _ := datastore.GetUserLogin(dbConn, email, password)
	if email == res.Email && password == res.Password {
		fmt.Println(res.Nama, "Berhasi Login")
		fmt.Println("")
		id := res.ID
		for input != 9 {
			fmt.Println("\t=========== DASHBOARD ===========")
			fmt.Println("\t]> Selamat Datang,", res.Nama)
			fmt.Println("\t---------------------------------")
			fmt.Println("\t1.\tUpdate Profil")
			fmt.Println("\t2.\tTambah Buku")
			fmt.Println("\t3.\tUpdate Buku")
			fmt.Println("\t4.\tDelete Buku")
			fmt.Println("\t5.\tDaftar Buku Saya")
			fmt.Println("\t6.\tPinjam Buku")
			fmt.Println("\t7.\tKembalikan Buku")
			fmt.Println("\t8.\tLogout")
			fmt.Println("\t= = = = = = = = = = = = = = = = =")
			fmt.Print("\tMasukkan pilihan menu :")
			fmt.Scanln(&input)
			if input == 1 {
				var nama string
				fmt.Print("Ganti Nama : ")
				fmt.Scanln(&nama)
				updateNama, err := datastore.UpdateUser(dbConn, nama, id)
				if err != nil {
					fmt.Println("terjadi sebuah kesalahan :", err)
				}
				res.Nama = updateNama.Nama
			} else if input == 2 {
				var newBook schema.Book

				fmt.Print("Masukkan Judul Buku:")
				fmt.Scanln(&newBook.Judul)
				fmt.Print("Masukkan Penerbit :")
				fmt.Scanln(&newBook.Penerbit)
				fmt.Print("Masukkan Penulis :")
				fmt.Scanln(&newBook.Penulis)
				fmt.Print("Masukkan tahun terbit :")
				fmt.Scanln(&newBook.TahunTerbit)

				_, err := datastore.InsertBuku(dbConn, newBook)
				if err != nil {
					fmt.Println("terjadi sebuah kesalahan :", err)
				}
				fmt.Println("Buku berhasil ditambahkan")
			} else if input == 3 {
				var newBook schema.Book
				fmt.Println("Ganti Judul : ")
				fmt.Scanln(&newBook.Judul)
				_, err := datastore.UpdateBuku(dbConn, newBook)
				if err != nil {
					fmt.Println("terjadi sebuah kesalahan :", err)
				}
				fmt.Println("Buku berhasil diupdate")
			} else if input == 4 {
				var newBook schema.Book
				fmt.Println("Hapus Buku : ")
				fmt.Scanln(&newBook.Judul)
				_, err := datastore.DeleteBuku(dbConn, newBook)
				if err != nil {
					fmt.Println("terjadi sebuah kesalahan :", err)
				}
				fmt.Println("Buku berhasil dihapus")
			}
		}
	} else {
		fmt.Println("email dan password salah")
		return
	}
}

func SelectBuku(dbConn *gorm.DB) {
	tampilBuku, err := datastore.TampilkanBuku(dbConn)
	if err != nil {
		fmt.Println("terjadi sebuah kesalahan :", err)
	}
	fmt.Println(tampilBuku)
}

func main() {

	// load .env as config database credentials
	config := ReadEnv()

	// make connection to mysql database
	dbConn := ConnectDB(config)

	// membuat database migration
	// dbConn.AutoMigrate(
	// 	&schema.User{},
	// 	&schema.Book{},
	// 	&schema.Transactions{},
	// )

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
			SelectBuku(dbConn)
		case 99:
			fmt.Println("Terimakasih banyak sudah mencoba program kami.")
			return
		default:
			fmt.Println("Maaf, pilihan inputan menu tidak tersedia!")
		} //end switch
	} //end loop
} //end main
