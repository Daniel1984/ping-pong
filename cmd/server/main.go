package main

import (
	"log"

	"github.com/ping-pong/pkg/hub"
	"github.com/ping-pong/pkg/sigtermhandler"
)

func main() {
	h, err := hub.Start()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("server ready to accept connections...")

	sigtermhandler.Init(func() {
		log.Print("shutting down the HUB")
		h.Shutdown()
	})
}
