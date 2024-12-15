package dto

import "github.com/Bazhenator/buffer/internal/entities"

type AppendRequestIn struct {
	Request *entities.Request
}

type PopTopOut struct {
	Request *entities.Request
}