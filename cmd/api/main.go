package main

import (
	"log"

	"github.com/phildehovre/go-template/cmd/api/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}
	log.Printf(`Server running on port %s`, app.config.addr)
	mux := app.mount()

	log.Fatal(app.run(mux))

}
