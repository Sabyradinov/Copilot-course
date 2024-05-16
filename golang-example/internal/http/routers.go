package http

func (s *server) routers() {

	s.engine.GET("/health")
}
