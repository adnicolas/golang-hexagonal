package main

import (
	"log"

	"github.com/adnicolas/golang-hexagonal/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
