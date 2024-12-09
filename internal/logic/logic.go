package logic

import (
	"context"

	"github.com/Bazhenator/buffer/internal/logic/dto"
)

type Logic struct {
}

func NewLogic() *Logic {
	return &Logic{}
}

func (l *Logic) AppendRequest(ctx context.Context, in *dto.AppendRequestIn) (*dto.AppendRequestOut, error) {
	return nil, nil
}

func (l *Logic) PopTop(ctx context.Context) (*dto.PopTopOut, error) {
	return nil, nil
}

func (l *Logic) PopBottom(ctx context.Context, in *dto.PopBottomIn) (*dto.PopBottomOut, error) {
	return nil, nil
}
