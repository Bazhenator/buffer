package delivery

import "github.com/Bazhenator/buffer/internal/entities"

type AppendRequestIn struct {
	Request *entities.Request
}

type AppendRequestOut struct {
	Size   uint64
	Status bool
}

type PopTopOut struct {
	Request *entities.Request
}

type PopBottomIn struct {
	CurRequest *entities.Request
}

type PopBottomOut struct {
	DeclinedRequest *entities.Request
	Status          bool
}