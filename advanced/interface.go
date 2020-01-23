package advanced

import "time"

// Option configures a Server.
type Option interface {
	apply(*Server)
}

// Timeout configures a maximum length of idle connection in Server.
type Timeout time.Duration

func (t Timeout) apply(s *Server) {
	s.timeout = time.Duration(t)
}

// MaxConnections configures a maximum number of active connectioons in Server.
type MaxConnections int

func (m MaxConnections) apply(s *Server) {
	s.maxConns = int(m)
}

// Server serves as the server part of a theoretical network application.
type Server struct {
	addr string

	timeout  time.Duration
	maxConns int
}

// NewServer initializes a new Server listening on addr with optional configuration.
func NewServer(addr string, opts ...Option) *Server {
	server := &Server{
		addr: addr,
	}

	for _, opt := range opts {
		opt.apply(server)
	}

	return server
}
