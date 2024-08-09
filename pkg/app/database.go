package app

import (
	"database/sql"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/ssss-tantalum/gophatt/ent"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewClient(dbDsn string) *ent.Client {
	db, err := sql.Open("pgx", dbDsn)
	if err != nil {
		log.Fatal(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)

	return ent.NewClient(ent.Driver(drv))
}
