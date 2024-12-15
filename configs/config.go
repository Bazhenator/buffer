package configs

import (
	"errors"
	"os"
	"strconv"

	"go.uber.org/multierr"

	"github.com/Bazhenator/tools/src/logger"
	grpcListener "github.com/Bazhenator/tools/src/server/grpc/listener"
)

const (
	EnvBufferCapacity = "BUFFER_CAPACITY"
)

// Config is a main configuration struct for application
type Config struct {
	Environment    string
	Grpc           *grpcListener.GrpcConfig
	BufferCapacity uint64
	LoggerConfig   *logger.LoggerConfig
}

// NewConfig returns application config instance
func NewConfig() (*Config, error) {
	var errorBuilder error

	grpcConfig, err := grpcListener.NewStandardGrpcConfig()
	multierr.AppendInto(&errorBuilder, err)

	loggerConfig, err := logger.NewLoggerConfig()
	multierr.AppendInto(&errorBuilder, err)

	EnvBufferCapacityStr, ok := os.LookupEnv(EnvBufferCapacity)
	if !ok {
		multierr.AppendInto(&errorBuilder, errors.New("BUFFER_CAPACITY is not defined"))
	}

	bufCapacity, err := strconv.Atoi(EnvBufferCapacityStr)
	multierr.AppendInto(&errorBuilder, err)

	if errorBuilder != nil {
		return nil, errorBuilder
	}

	glCfg := &Config{
		Grpc:           grpcConfig,
		BufferCapacity: uint64(bufCapacity),
		LoggerConfig:   loggerConfig,
	}

	return glCfg, nil
}
