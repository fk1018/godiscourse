package main

import (
	"context"
	"log"

	"github.com/godiscourse/godiscourse/api/durable"
	_ "github.com/lib/pq"
)

func main() {
	db := durable.OpenDatabaseClient(context.Background())
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Panicln(err)
	}

	if err := startHTTP(db); err != nil {
		log.Panicln(err)
	}
}
