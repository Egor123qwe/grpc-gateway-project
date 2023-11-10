package storage

import "github.com/Egor123qwe/grpc-gateway-project/internal/storage/repsInterfaces"

type Storage interface {
	User() repsInterfaces.User
}
