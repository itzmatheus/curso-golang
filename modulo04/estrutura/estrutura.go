package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	psqlInfo := "host=127.0.0.1 port=32769 dbname=cursogo user=postgres password=postgres sslmode=disable" // criar banco de dados cursogo

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, `DROP TABLE IF EXISTS usuarios;`)
	exec(db, ` CREATE TABLE usuarios(
		id SERIAL,
		nome character varying(256),
		primary key(id)
	)`)

}
