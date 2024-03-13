package rest

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type server struct {
	srv    *http.Server
	logger *logrus.Entry
}

func NewServer(host, port string, log *logrus.Logger) *server {
	return &server{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", host, port),
			Handler: getRouter(log),
		},
		logger: log.WithField("service", "rest"),
	}
}

func (s *server) Run() error {
	const op = "rest.server.Run"

	s.logger.Info("REST server is running")
	if err := s.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *server) Stop() error {
	s.logger.Info("REST server is stopping")
	if err := s.srv.Shutdown(nil); err != nil {
		return fmt.Errorf("rest.server.Stop: %w", err)
	}
	return nil
}
