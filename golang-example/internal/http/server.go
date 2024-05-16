package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Sabyradinov/Copilot-course/golang-example/config"
	"github.com/gin-gonic/gin"
)

type IServer interface {
	Start() (sc chan error)
	Stop() error
}

type server struct {
	engine *gin.Engine
	srv    *http.Server
	srvCh  chan error
}

func NewServer(cfg config.Config) (s IServer, err error) {

	engine := gin.New()

	srv := &server{
		engine: engine,
		srv: &http.Server{
			Addr:    ":" + strconv.Itoa(cfg.WebServer.Port),
			Handler: engine,
		},
	}

	srv.routers()

	return srv, nil
}

func (s *server) Start() (sc chan error) {
	s.srvCh = make(chan error)

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.srvCh <- err
		}
	}()

	return s.srvCh
}

func (s *server) Stop() error {
	ctx := context.Background()

	return s.srv.Shutdown(ctx)
}
