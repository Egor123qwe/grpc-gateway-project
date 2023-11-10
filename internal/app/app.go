package app

import (
	"github.com/Egor123qwe/grpc-gateway-project/internal/app/server"
	"github.com/joho/godotenv"
)

func Start() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := server.New().Start(); err != nil {
		return err
	}

	return nil
}
