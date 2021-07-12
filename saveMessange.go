package main

import (
	"database/sql"
	"fmt"
	"log"
)

func SaveMessage (message, name string){
	DB, err := sql.Open("sqlite3", "telecart")
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	result, err := DB.Exec("insert into messages (message,name) values ( $1, $2)",
		message, name)
	if err != nil{
		log.Fatal(err)
	}
	msgId, err :=result.LastInsertId()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("\nThe message has been saved, id = ",msgId)

}