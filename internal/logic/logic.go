package logic

import (
	"context"

	"github.com/Bazhenator/buffer/configs"
	"github.com/Bazhenator/buffer/internal/entities"
	"github.com/Bazhenator/buffer/internal/logic/dto"
	"github.com/Bazhenator/tools/src/logger"
)

type Logic struct {
	c *configs.Config
	l *logger.Logger

	buffer *entities.Buffer
}

func NewLogic(c *configs.Config, l *logger.Logger, b *entities.Buffer) *Logic {
	return &Logic{
		c: c,
		l: l,

		buffer: b,
	}
}

func (l *Logic) AppendRequest(ctx context.Context, in *dto.AppendRequestIn) (*dto.AppendRequestOut, error) {
	l.l.InfoCtx(ctx, "AppendRequest started with", logger.NewField("data", in))

	// Check available space in l.buffer
	if l.buffer.GetSize() == l.buffer.GetCapacity() {
		l.l.Info("buffer overfilled", logger.NewField("size", l.buffer.GetSize()))

		// Getting previous request from l.buffer
		prevReq := l.buffer.GetTail()

		curPriority := in.Request.Priority
		prevPriority := prevReq.Priority

		// Check if priority of l.buffer.Tail less than the priority of in.Request
		if prevPriority < curPriority {
			declinedReq, err := l.buffer.PopBottom()
			if err != nil {
				l.l.ErrorCtx(ctx, "PopBottom failed with error:", logger.NewErrorField(err))
				return nil, err
			}
			l.l.InfoCtx(ctx, "PopBottom finished successfully with", logger.NewField("declined", declinedReq))

			_ = l.buffer.Append(in.Request)
			l.l.InfoCtx(ctx, "AppendRequest finished successfully with", logger.NewField("data", in))
		} else {
			l.l.InfoCtx(ctx, "AppendRequest failed with low priority", logger.NewField("priority", in.Request.Priority))
			l.l.InfoCtx(ctx, "Request declined", logger.NewField("id", in.Request.Id))
			return &dto.AppendRequestOut{Size: l.buffer.GetSize(), Status: false}, nil
		}
	} else {
		l.l.Info("buffer has available space", logger.NewField("size", l.buffer.GetSize()))
		
		_ = l.buffer.Append(in.Request)
		l.l.InfoCtx(ctx, "AppendRequest finished successfully with", logger.NewField("data", in))
	}

	return &dto.AppendRequestOut{Size: l.buffer.GetSize(), Status: true}, nil
}

func (l *Logic) PopTop(ctx context.Context) (*dto.PopTopOut, error) {
	l.l.InfoCtx(ctx, "PopTop started with", logger.NewField("head", l.buffer.GetHead()))

	req, err := l.buffer.PopTop()
	if err != nil {
		l.l.ErrorCtx(ctx, "PopTop failed with error:", logger.NewErrorField(err))
		return nil, err
	}

	l.l.InfoCtx(ctx, "PopTop finished successfully with", logger.NewField("request", req))
	return &dto.PopTopOut{Request: req}, nil
}