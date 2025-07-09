package main

import (
	"fmt"
	"go-mysql/database"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("MySql and Go")
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Conexion a la base de datos MySql Exitosa...!!!")

	/* dns := "jdanielmq:JdmQ1481@tcp(localhost:3306)/db_grossgym"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}*/

}
