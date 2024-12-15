package delivery

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Bazhenator/buffer/configs"
	"github.com/Bazhenator/buffer/internal/entities"
	"github.com/Bazhenator/buffer/internal/logic/dto"
	buffer "github.com/Bazhenator/buffer/pkg/api/grpc"
	"github.com/Bazhenator/tools/src/logger"
)

type BufferServer struct {
	buffer.UnimplementedBufferServiceServer

	c *configs.Config
	l *logger.Logger

	logic BufferLogic
}

func NewBufferServer(c *configs.Config, l *logger.Logger, logic BufferLogic) *BufferServer {
	return &BufferServer{
		c: c,
		l: l,

		logic: logic,
	}
}

func (s *BufferServer) AppendRequest(ctx context.Context, in *buffer.AppendRequestIn) (*emptypb.Empty, error) {
	s.l.DebugCtx(ctx, "AppendRequest started with", logger.NewField("data", in))
	req := in.GetReq()

	err := s.logic.AppendRequest(ctx, &dto.AppendRequestIn{
		Request: &entities.Request{
			Id:           req.GetId(),
			ClientId:     req.GetClientId(),
			CleaningType: entities.CleaningType(req.GetCleaningType()),
			Priority:     entities.Priority(req.GetPriority()),
			GeneratorId:  *req.GeneratorId,
		},
	})
	if err != nil {
		s.l.ErrorCtx(ctx, "AppendRequest error", logger.NewErrorField(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *BufferServer) PopTop(ctx context.Context, _ *emptypb.Empty) (*buffer.PopTopOut, error) {
	s.l.DebugCtx(ctx, "PopTop data", logger.NewField("in", nil))

	answer, err := s.logic.PopTop(ctx)
	if err != nil {
		s.l.ErrorCtx(ctx, "PopTop error", logger.NewErrorField(err))
		return nil, err
	}

	totalTime := answer.Request.TimeInBuffer.Seconds()

	return &buffer.PopTopOut{
		Req: &buffer.Request{
			Id:           answer.Request.Id,
			ClientId:     answer.Request.ClientId,
			CleaningType: uint32(answer.Request.CleaningType),
			Priority:     uint32(answer.Request.Priority),
			Status:       &answer.Request.Status,
			GeneratorId:  &answer.Request.GeneratorId,
			TimeInBuffer: &totalTime,
		},
	}, nil
}
