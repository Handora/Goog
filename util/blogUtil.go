package util

import (
	"database/sql"
	"os"
	"encoding/json"
	"fmt"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type Configuration struct {
	DataBase string
	Host string
	Port string
	User string
	Password string
}

type Article struct {
	Id      int
	Title   string
	Content string
	Time    time.Time
}

type JsonErr struct {
	Code int
	Text string
}

type GetRequest struct {
	Title   string
	Content string
	Tag     []int
}

type PostRequest struct {
	Id 		int
	Title   string
	Content string
	Tag     []int
}

type JsonOk struct {
	Code int
	Text string
}

type GetOk struct {
	Code int
	Text string
	Body interface{}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func DbInit() {
	file, err := os.Open("config.json")
	CheckError(err)
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	CheckError(err)

	dbLogin := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.User,  config.Password, config.Host, config.Port, config.DataBase)
	// fmt.Println(dbLogin)

	Db, err = sql.Open("mysql", dbLogin)
	CheckError(err)
}
