package main

import (
	"MyWebProject/dbInterface"
	"MyWebProject/handlers"
	"MyWebProject/handlers/user"
	"database/sql"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var UserObj dbInterface.UserService
var TaskObj dbInterface.TaskService
var statusObj dbInterface.StatusService

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
	TaskObj = &dbInterface.DbRepo{SqlConnection: DBConn}

	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", getUserHandler) // http://localhost:8800/user/sdvkjrni3f3
	router.HandleFunc("/user/create", createUserHandler)

	router.HandleFunc("/task/{id}", handlers.Taskhandler)

	http.ListenAndServe(":8800", router)

}

func getUserHandler(resp http.ResponseWriter, req *http.Request) {
	user.GetUser(resp, user.GetUserRequestWrapper{
		Req:     req,
		UserSvc: UserObj,
	})
}

func createUserHandler(resp http.ResponseWriter, req *http.Request) {
	user.CreateUser(resp, user.CreateUserRequestWrapper{
		Req:     req,
		UserSvc: UserObj,
	})
}
