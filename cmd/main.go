package main

import (
	"github.com/Egor123qwe/grpc-gateway-project/internal/app"
	"log"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
