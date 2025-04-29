package main

import (
	"Amani_Classes/MyWebProject/dbInterface"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/exampledb")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	db, err := connectToDB()
	if err != nil {
		panic(err)
	}

	var userCOnfusion dbInterface.UserService

	userCOnfusion = &dbInterface.DbRepo{DB: db}
	// userCOnfusion.GetUser("User1")

	userCOnfusion = &dbInterface.TaskRepo{DB: db}

	// userCOnfusion.GetUser("User2")
	// userRepoObj.GetUser("user1")

	// user, err := intf.GetUser("")
	// if err != nil {
	// 	//return error
	// }
	//process user
}
