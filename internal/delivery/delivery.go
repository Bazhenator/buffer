package delivery

import (
	"context"

	"github.com/Bazhenator/buffer/configs"
	"github.com/Bazhenator/buffer/internal/entities"
	"github.com/Bazhenator/buffer/internal/logic"
	"github.com/Bazhenator/buffer/pkg/api/grpc"
	"github.com/Bazhenator/tools/src/logger"
	"github.com/Bazhenator/buffer/internal/logic/dto"
)

type BufferServer struct {
	buffer.UnimplementedBufferServiceServer

	c *configs.Config
	l *logger.Logger
	
	logic logic.BufferLogic
}

func NewBufferServer(c *configs.Config, l *logger.Logger, logic logic.BufferLogic) *BufferServer {
	return &BufferServer{
		c: c,
		l: l,

		logic: logic,
	}
} 

func (s *BufferServer) AppendRequest(ctx context.Context, in *buffer.AppendRequestIn) (*buffer.AppendRequestOut, error) {
	s.l.DebugCtx(ctx, "AppendRequest data", logger.NewField("in", in))
	req := in.GetReq()

	answer, err := s.logic.AppendRequest(ctx, &dto.AppendRequestIn{
		Request: &entities.Request {
			Id: req.GetId(),
			ClientId: req.GetClientId(),
			CleaningType: entities.CleaningType(req.GetCleaningType()),
			Priority: entities.Priority(req.GetPriority()),
		},
	})
	if err != nil {
		s.l.ErrorCtx(ctx, "AppendRequest error", logger.NewErrorField(err))
		return nil, err
	}

	return &buffer.AppendRequestOut{
		Size: answer.Size,
		Status: answer.Status,
	}, nil
}

func (s *BufferServer) PopTop(ctx context.Context, _ *buffer.PopTopIn) (*buffer.PopTopOut, error) {
	s.l.DebugCtx(ctx, "PopTop data", logger.NewField("in", nil))

	answer, err := s.logic.PopTop(ctx)
	if err != nil {
		s.l.ErrorCtx(ctx, "PopTop error", logger.NewErrorField(err))
		return nil, err
	}

	return &buffer.PopTopOut{
		Req: &buffer.Request{
			Id: answer.Request.Id,
			ClientId: answer.Request.ClientId,
			CleaningType: uint32(answer.Request.CleaningType),
			Priority: uint32(answer.Request.Priority),
		},
	}, nil
}