package main

import (
	"github.com/edreg/awesome/app/server"
	"log"
	"net/http"
)

func main() {
	pServer := &server.PlayerServer{Store: server.NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", pServer); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
