package main

import (
	"ProjectIdeas/monolith/api"
	"ProjectIdeas/monolith/httpserver"
	"ProjectIdeas/monolith/internal/db"
	"ProjectIdeas/monolith/internal/db/sqlite"
	"context"
	"log"
	"net/http"
	"time"
)

const port = "8080" // 8080 is the port number for the monolith service

func main() {
	ctx := context.Background()

	var dber db.DBer
	for {
		dberr, err := sqlite.New(ctx)
		if err != nil {
			log.Printf("db not ready. error : `%s`. Retrying in 2 seconds...\n", err)
			time.Sleep(time.Second * 4)
			continue
		}
		dber = dberr
		break
	}

	a, err := api.New(api.ApiOptions{DB: dber})
	if err != nil {
		log.Println("error creating new api:", err)
	}

	log.Printf("Starting monolith service on port: %s", port)

	if err := http.ListenAndServe("localhost:8080", httpserver.NewRouter(ctx, httpserver.RouterOptions{Api: a})); err != nil {
		log.Fatalf("fatal crash %s", err)
	}
}
