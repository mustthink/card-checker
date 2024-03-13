package internal

import (
	"flag"

	"github.com/sirupsen/logrus"

	"github.com/mustthink/card-checker/internal/grpc"
	"github.com/mustthink/card-checker/internal/rest"
)

type Server interface {
	Run() error
	Stop() error
}

const (
	restType = "rest"
	grpcType = "grpc"
)

func NewServer() Server {
	var serverType, host, port string
	flag.StringVar(&serverType, "s", restType, "server type")
	flag.StringVar(&host, "h", "", "server host")
	flag.StringVar(&port, "p", "8081", "server port")
	flag.Parse()

	logger := logrus.New()
	switch serverType {
	case restType:
		return rest.NewServer(host, port, logger)

	case grpcType:
		return grpc.NewServer(port, logger)

	default:
		logger.Fatal("unknown server type")
	}

	return nil
}
