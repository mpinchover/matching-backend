package messaging

import "matching/src/types/entities"

type IMessaging interface {
	CreateRoom(room *entities.Room) (*entities.Room, error)
}

type Messaging struct {
}

func New() *Messaging {
	return &Messaging{}
}

func (m *Messaging) CreateRoom(room *entities.Room) (*entities.Room, error) {
	return nil, nil
}
