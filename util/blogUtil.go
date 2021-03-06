package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"net/rpc"
)

// global db for all sql usage
var Db *sql.DB

// config.json structure
type Configuration struct {
	DataBase string
	Host     string
	Port     string
	User     string
	Password string
}

// article structure
type Article struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	Intro   string    `json:"intro"`
	Content string    `json:"content"`
	Tag     []Tag     `json:"tag"`
	Time    time.Time `json:"time"`
}

// paged articles structure for json
type PagedArticles struct {
	First  Article `json:"0"`
	Second Article `json:"1"`
	Third  Article `json:"2"`
	Fouth  Article `json:"3"`
	Fifth  Article `json:"4"`
}

// tag structure
type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// timeline structure
type Timeline struct {
	Year   int     `json:"year"`
	Events []Event `json:"events"`
}

type Event struct {
	Id     int       `json:"id"`
	Months int       `json:"months"`
	Times  time.Time `json:"times"`
	Even   string    `json:"even"`
}

// unified structure for response structure
type Response struct {
	Code int
	Text string
	Body interface{}
}

type ArticlesResponse struct {
	Code        int
	Text        string
	Body        interface{}
	Total       int `json:"total"`
	CurrentPage int `json:"currentPage"`
	Length      int `json:"length"`
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.WriteHeader(http.StatusNotFound)

		if err := json.NewEncoder(w).Encode(Response{Code: statusCode, Text: message}); err != nil {
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
	dbLogin := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.User, config.Password, config.Host, config.Port, config.DataBase)
	// fmt.Println(dbLogin)

	Db, err = sql.Open("mysql", dbLogin)
	CheckError(err)
}
