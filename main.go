package main

import (
	"MyWebProject/dbInterface"
	"database/sql"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var UserObj dbInterface.UserService

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/taskmanager")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	DBConn, err := connectToDB()
	if err != nil {
		panic(err)
	}

	UserObj = &dbInterface.DbRepo{SqlConnection: DBConn}

	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", GetUser) // http://localhost:8800/user/sdvkjrni3f3

	http.ListenAndServe(":8800", router)

}
