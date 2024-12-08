package delivery

import (
	"context"

	"github.com/Bazhenator/buffer/configs"
	"github.com/Bazhenator/buffer/pkg/api/grpc"
	"github.com/Bazhenator/tools/src/logger"
)

type BufferServer struct {
	buffer.UnimplementedBufferServiceServer

	c *configs.Config
	l *logger.Logger
	
	// logic logic.BufferLogic
}

// TODO: adding logic initializing
func NewBufferServer(c *configs.Config, l *logger.Logger) *BufferServer {
	return &BufferServer{
		c: c,
		l: l,

		// logic: logic
	}
} 

func (s *BufferServer) AppendRequest(ctx context.Context, in *buffer.AppendRequestIn) (*buffer.AppendRequestOut, error) {
	return nil, nil
}

func (s *BufferServer) PopTop(ctx context.Context, in *buffer.PopTopIn) (*buffer.PopTopOut, error) {
	return nil, nil
}

func (s *BufferServer) PopBottom(ctx context.Context, in *buffer.PopBottomIn) (*buffer.PopBottomOut, error) {
	return nil, nil
}