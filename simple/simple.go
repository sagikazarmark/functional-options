package simple

import "time"

// Option configures a Server.
type Option func(*Server)

// Timeout configures a maximum length of idle connection in Server.
func Timeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// MaxConnections configures a maximum number of active connectioons in Server.
func MaxConnections(maxConns int) func(*Server) {
	return func(s *Server) {
		s.maxConns = maxConns
	}
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
		opt(server)
	}

	return server
}
