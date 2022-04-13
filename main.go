package main

import (
	"fmt"
	"os"
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
	res.Username = os.Getenv("USERNAME")
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

	config := ReadEnv()

	conn := ConnectDB(config)

	fmt.Println(conn)
	fmt.Println(conn.Error)

}
