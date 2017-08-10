package util

import (
	"database/sql"
	"os"
	"encoding/json"
	"fmt"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

// global db for all sql usage
var Db *sql.DB

// config.json structure
type Configuration struct {
	DataBase string
	Host string
	Port string
	User string
	Password string
}

// article structure
type Article struct {
	Id      int
	Title   string
	Content string
	Time    time.Time
}

// tag structure
type Tag struct {
	Id int
	Name string
}

// unified structure for response structure
type Response struct {
	Code int
	Text string
	Body interface{}
}

// just check error, and if error, we panic
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckAndResponse(w http.ResponseWriter, err error, statusCode int, message string) {
	if err != nil {
		// not found, set the header, and write the statusNotFound with body
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)

		if err := json.NewEncoder(w).Encode(Response{Code:statusCode, Text: message}); err != nil {
			panic(err)
		}
		panic(err)
	}
}

func DbInit(conf string) {
	// open the file, and read it, must be json
	file, err := os.Open(conf)
	CheckError(err)

	// decode the json
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	CheckError(err)

	// use the mysql driver to connect mysql server in remote
	dbLogin := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.User,  config.Password, config.Host, config.Port, config.DataBase)
	// fmt.Println(dbLogin)

	Db, err = sql.Open("mysql", dbLogin)
	CheckError(err)
}
