package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	psqlInfo := "host=127.0.0.1 port=32769 dbname=cursogo user=postgres password=postgres sslmode=disable" // criar banco de dados cursogo

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE usuarios SET nome = 'Matheus Leite' WHERE id = 13")
	if err != nil {
		log.Panic(err)
	}

}
