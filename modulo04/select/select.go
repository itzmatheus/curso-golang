package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type usuario struct {
	id   int
	nome string
}

func main() {

	psqlInfo := "host=127.0.0.1 port=32769 dbname=cursogo user=postgres password=postgres sslmode=disable" // criar banco de dados cursogo

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id > $1", 1)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}

}
