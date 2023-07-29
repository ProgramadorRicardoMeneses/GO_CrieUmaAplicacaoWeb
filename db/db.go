package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBD() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=ricardo host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err)
	}
	return db

}
