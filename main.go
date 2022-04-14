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
	Port     int64
	DB       string
}

func ConnectDB(configData Config) *gorm.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		configData.Username,
		configData.Password,
		configData.Host,
		configData.Port,
		configData.DB)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Println("Koneksi gagal", err)
		return nil
	}
	return db
}

func ReadEnv() Config {
	if err := godotenv.Load("local.env"); err != nil {
		fmt.Println("Error load file", err)
	}
	res := Config{}
	res.Username = os.Getenv("User")
	res.DB = os.Getenv("DB")
	res.Password = os.Getenv("Password")
	res.Host = os.Getenv("Host")
	intConv, _ := strconv.Atoi(os.Getenv("Port"))
	res.Port = int64(intConv)
	return res
}

func main() {

	config := ReadEnv()

	conn := ConnectDB(config)

	fmt.Println(conn)
	fmt.Println(conn.Error)

	func Register(nama,hp,alamat,email,password string){
		userAcc := datastore.UserDB{Db: conn}
		res,err = userAcc.RegisterUser(schema.User{Nama: &nama, Hp: &hp, Alamat: &alamat,Email: &email, Password: &password})
		return res.Ro 
	}


}
