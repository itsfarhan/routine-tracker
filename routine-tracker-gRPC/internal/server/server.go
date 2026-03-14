// Package server wires the gRPC handlers to the domain layer.
package server

import (
	"routine-tracker/routine-tracker-gRPC/api"
	"strconv"

	"fmt"
	"net"

	"google.golang.org/grpc"
)

// Server is the implemetation of the gRPC server.

type Server struct {
	api.UnimplementedHabitsServer
	lgr Logger
}

// New creates a new Server that can Listen and Serve gRPC requests.
func New(lgr Logger) *Server {
	return &Server{
		lgr: lgr,
	}
}

type Logger interface {
	Logf(format string, args ...any)
}

func (s *Server) ListenAndServe(port int) error {
	const addr = "127.0.0.1"
	listener, err := net.Listen("tcp", net.JoinHostPort(addr, strconv.Itoa(port)))
	if err != nil {
		return fmt.Errorf("cannot listen on %s:%d: %w", addr, port, err)
	}
	// creates grpc server and wires your Server to it
	grpcServer := grpc.NewServer()
	api.RegisterHabitsServer(grpcServer, s) // wires your Server to gRPC
	s.lgr.Logf("starting server on port %d\n", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("error while listening %w", err)
	}
	return grpcServer.Serve(listener)
}

// server.go
type Repository interface {
    Add(ctx context.Context, habit habit.Habit) error
    FindAll(ctx context.Context) ([]habit.Habit, error)
}

type Server struct {
    api.UnimplementedHabitsServer
    db  Repository
    lgr Logger
}

func New(repo Repository, lgr Logger) *Server {
    return &Server{db: repo, lgr: lgr}
}