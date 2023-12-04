package pgdb

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "postgres://postgres:F*bEc6CDaaB23aCb3ED-ad-g*4dA4CFG@roundhouse.proxy.rlwy.net:19701/railway")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
