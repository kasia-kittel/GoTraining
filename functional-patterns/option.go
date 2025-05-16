package patterns

import "time"

// Functional Options
type Server struct {
	host    string
	port    int
	timeOut time.Duration
}

type Option func(*Server)

func WithHost(h string) Option {
	return func(s *Server) { s.host = h }
}

func WithPort(p int) Option {
	return func(s *Server) { s.port = p }
}

func WithTimeout(t time.Duration) Option {
	return func(s *Server) { s.timeOut = t }
}

func NewServer(opt ...Option) *Server {
	server := &Server{}
	for _, v := range opt {
		v(server)
	}
	return server
}
