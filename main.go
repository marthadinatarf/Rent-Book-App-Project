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
		user_id := res.ID
		for input != 9 {
			fmt.Println("=========== DASHBOARD ===========")
			fmt.Println("]> Selamat Datang,", res.Nama)
			fmt.Println("---------------------------------")
			fmt.Println("1.\tUpdate Profil")
			fmt.Println("2.\tTambah Buku")
			fmt.Println("3.\tUpdate Buku")
			fmt.Println("4.\tDelete Buku")
			fmt.Println("5.\tPinjam Buku")
			fmt.Println("6.\tKembalikan Buku")
			fmt.Println("9.\tLogout")
			fmt.Println("= = = = = = = = = = = = = = = = =")
			fmt.Print("Masukkan pilihan menu :")
			fmt.Scanln(&input)
			if input == 1 {
				var nama string
				fmt.Print("Ganti Nama : ")
				fmt.Scanln(&nama)
				updateNama, err := datastore.UpdateUser(dbConn, nama, user_id)
				if err != nil {
					fmt.Println("terjadi sebuah kesalahan :", err)
				}
				res.Nama = updateNama.Nama
			} else if input == 2 {
				var books schema.Book
				books.UserID = user_id
				fmt.Print("Masukkan Judul Buku:")
				fmt.Scanln(&books.Judul)
				fmt.Print("Masukkan Penerbit :")
				fmt.Scanln(&books.Penerbit)
				fmt.Print("Masukkan Penulis :")
				fmt.Scanln(&books.Penulis)
				fmt.Print("Masukkan tahun terbit :")
				fmt.Scanln(&books.TahunTerbit)

				_, err := datastore.InsertBuku(dbConn, schema.Book{
					Judul:       books.Judul,
					Penerbit:    books.Penerbit,
					Penulis:     books.Penulis,
					TahunTerbit: books.TahunTerbit,
					UserID:      books.UserID,
				})
				if err != nil {
					fmt.Println("Terjadi kesalahan saat tambah Buku :", err)
				}
				fmt.Println("Buku berhasil ditambahkan")
			} else if input == 3 {
				var books schema.Book
				fmt.Print("Pilih id Buku : ")
				fmt.Scanln(&books.ID)
				book, _ := datastore.GetBukuById(dbConn, books.ID)
				fmt.Println("-+-+-+-+-+-+-+-+-+-+-+")
				fmt.Printf("| Judul\t: %s \n", book.Judul)
				fmt.Printf("| Penerbit\t: %s \n", book.Penerbit)
				fmt.Printf("| Penulis\t: %s \n", book.Penulis)
				fmt.Printf("| Release\t: %d \n", book.TahunTerbit)
				fmt.Println("-+-+-+-+-+-+-+-+-+-+-+")
				fmt.Print("Edit Judul : ")
				fmt.Scanln(&books.Judul)
				fmt.Print("Edit Penerbit : ")
				fmt.Scanln(&books.Penerbit)
				fmt.Print("Edit Penulis : ")
				fmt.Scanln(&books.Penulis)
				fmt.Print("Edit Tahun Terbit : ")
				fmt.Scanln(&books.TahunTerbit)
				_, err := datastore.UpdateBuku(dbConn, books.ID, schema.Book{
					Judul:       books.Judul,
					Penerbit:    books.Penerbit,
					Penulis:     books.Penulis,
					TahunTerbit: books.TahunTerbit,
				})
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
			} else if input == 5 {
				//var transaction schema.Book
				fmt.Print("Masukkan tanggal peminjaman : ")
				//fmt.Scanln(&)
				fmt.Print("Masukkan Penerbit :")
				//fmt.Scanln(&books.Penerbit)
				fmt.Print("Masukkan Penulis :")
				//fmt.Scanln(&books.Penulis)
				fmt.Print("Masukkan tahun terbit :")
				//fmt.Scanln(&books.TahunTerbit)

				// _, err := datastore.InsertBuku(dbConn, schema.Book{
				// 	Judul:       books.Judul,
				// 	Penerbit:    books.Penerbit,
				// 	Penulis:     books.Penulis,
				// 	TahunTerbit: books.TahunTerbit,
				// 	UserID:      books.UserID,
				// })
				// if err != nil {
				// 	fmt.Println("Terjadi kesalahan saat tambah Buku :", err)
				// }
				fmt.Println("Buku berhasil ditambahkan")
			} else if input == 6 {
				//var transaction schema.Book
				fmt.Print("Masukkan tanggal peminjaman : ")
				//fmt.Scanln(&)
				fmt.Print("Masukkan Penerbit :")
				//fmt.Scanln(&books.Penerbit)
				fmt.Print("Masukkan Penulis :")
				//fmt.Scanln(&books.Penulis)
				fmt.Print("Masukkan tahun terbit :")
				//fmt.Scanln(&books.TahunTerbit)

				// _, err := datastore.InsertBuku(dbConn, schema.Book{
				// 	Judul:       books.Judul,
				// 	Penerbit:    books.Penerbit,
				// 	Penulis:     books.Penulis,
				// 	TahunTerbit: books.TahunTerbit,
				// 	UserID:      books.UserID,
				// })
				// if err != nil {
				// 	fmt.Println("Terjadi kesalahan saat tambah Buku :", err)
				// }
				fmt.Println("Buku berhasil ditambahkan")
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
		fmt.Println("4.\tLihat Detail Buku")
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
		case 4:
			var books schema.Book
			fmt.Print("Pilih id Buku : ")
			fmt.Scanln(&books.ID)
			book, _ := datastore.GetBukuById(dbConn, books.ID)
			fmt.Println("-+-+-+-+-+-+-+-+-+-+-+")
			fmt.Printf("| Judul\t\t: %s \n", book.Judul)
			fmt.Printf("| Penerbit\t: %s \n", book.Penerbit)
			fmt.Printf("| Penulis\t: %s \n", book.Penulis)
			fmt.Printf("| Release\t: %d \n", book.TahunTerbit)
			fmt.Println("-+-+-+-+-+-+-+-+-+-+-+")
		case 99:
			fmt.Println("Terimakasih banyak sudah mencoba program kami.")
			return
		default:
			fmt.Println("Maaf, pilihan inputan menu tidak tersedia!")
		} //end switch
	} //end loop
} //end main
