package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {

	psqlInfo := "host=127.0.0.1 port=32769 dbname=cursogo user=postgres password=postgres sslmode=disable" // criar banco de dados cursogo

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	tx.Exec("insert into usuarios(nome) values('Matheus')")
	tx.Exec("insert into usuarios(nome) values('José')")
	tx.Exec("insert into usuarios(nome) values('Leite')")
	// _, err = tx.Exec("insert into usuarios(id, nome) values(1, 'Mendes')") // chave duplicada
	// if err != nil {
	// 	tx.Rollback()
	// 	log.Fatal(err)
	// }
	tx.Commit()

}
