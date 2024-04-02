package main

import (
	"cloudrun/internal/infra/webserver"
	"log"
)

func main() {
	if err := webserver.StartServer(); err != nil {
		log.Panic(err)
	}
}
