package scenarios

import (
	"github.com/Egor123qwe/grpc-gateway-project/internal/scenarios/serviceInterfaces"
	"github.com/Egor123qwe/grpc-gateway-project/internal/storage"
)

type Scenarios struct {
	storage storage.Storage
}

func New(storage storage.Storage) serviceInterfaces.User {
	return &Scenarios{
		storage: storage,
	}
}

func (s *Scenarios) CreateUser() {

}

func (s *Scenarios) GetUser() {

}

func (s *Scenarios) DeleteUser() {

}

func (s *Scenarios) EditUser() {

}

func (s *Scenarios) SubscribeUser() {

}

func (s *Scenarios) UnsubscribeUser() {

}
