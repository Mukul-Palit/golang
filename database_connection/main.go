package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "USERNAME:PASSWORD@tcp(LOCAL OR IP_ADDRESS:PORT_NUMBER)/DATABASE_NAME")   
    if err != nil {
	//Print Error generated during the database connection
        panic(err.Error())
    }
	//Closing the db connection after all the operations
	defer db.Close()
	fmt.Println("successfully connected to the database")
	insert, err := db.Query("Insert into TABLE_NAME values (VALUE1,VALUE2)")
	if err != nil {
        panic(err.Error())
    }
	defer insert.Close()

}
