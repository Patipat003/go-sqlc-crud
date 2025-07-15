package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/patipat003/go-dbsql-2/controllers"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "postgresql://postgres:1234@localhost:5432/testdb?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// err = controllers.ReadUser(db);
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = controllers.ReadUserById(db, 7)
	if err != nil {
		log.Fatal(err)
	}

	// err = controllers.AddUser(db);  
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = controllers.DelUser(db, 22)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = controllers.UpdateDataUser(db, "LUCY", "LUCZ", 5)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
}