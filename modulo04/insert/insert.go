package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	psqlInfo := "host=127.0.0.1 port=32769 dbname=cursogo user=postgres password=postgres sslmode=disable" // criar banco de dados cursogo

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Erro abrindo db", err)
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into usuarios(nome) values('Maria')")
	if err != nil {
		log.Println("erro maria ", err)
	}
	_, err = db.Exec("insert into usuarios(nome) values('João')")
	if err != nil {
		log.Println("erro joao ", err)
	}

	res, err := db.Exec("insert into usuarios(nome) values('Pedro')")
	if err != nil {
		log.Println("erro pedro ", err)
	}

	id, _ := res.LastInsertId()

	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)

}
