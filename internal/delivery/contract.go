package delivery

import (
	"context"

	"github.com/Bazhenator/buffer/internal/logic/dto"
)

type BufferLogic interface {
	AppendRequest(context.Context, *dto.AppendRequestIn) error
	PopTop(context.Context) (*dto.PopTopOut, error)
}
