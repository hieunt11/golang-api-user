package connect

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Openconnect() *sql.DB {
	db, err := sql.Open("mysql", "hieunt20:Hanoi@123@tcp(127.0.0.1:3306)/user-db")
	if err != nil {
		log.Print(err.Error())
	}
	//defer db.Close()
	return db
}
