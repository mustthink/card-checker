package grpc

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	card_checker "github.com/mustthink/card-checker/internal/grpc/gen"
)

type server struct {
	server *grpc.Server
	port   string
	logger *logrus.Entry
}

func NewServer(port string, log *logrus.Logger) *server {
	grpcServer := grpc.NewServer()
	card_checker.RegisterValidatorServer(grpcServer, NewValidator())

	return &server{
		server: grpcServer,
		port:   port,
		logger: log.WithField("service", "grpc"),
	}
}

func (s *server) Run() error {
	const op = "grpc.server.Run"

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.logger.Info("gRPC server is running")
	if err := s.server.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *server) Stop() error {
	s.server.Stop()
	s.logger.Info("gRPC server is stopped")
	return nil
}
