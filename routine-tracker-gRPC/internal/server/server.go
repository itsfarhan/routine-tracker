package server

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/itsfarhan/routine-tracker/api"
	"github.com/itsfarhan/routine-tracker/internal/habit"
	"google.golang.org/grpc"
)

// Repository is the interface the server uses to talk to storage.
// Defined here so tests can inject a mock instead of the real repository.
type Repository interface {
	Add(ctx context.Context, habit habit.Habit) error
	FindAll(ctx context.Context) ([]habit.Habit, error)
	Find(ctx context.Context, id habit.ID) (habit.Habit, error)
	AddTick(ctx context.Context, id habit.ID, t time.Time) error
}

// Logger is the interface for logging inside the server.
// testing.T satisfies this, so you can pass t directly in tests.
type Logger interface {
	Logf(format string, args ...any)
}

// Server is the implementation of the gRPC server.
type Server struct {
	api.UnimplementedHabitsServer // required — keeps old code working when new RPCs are added
	db                            Repository
	lgr                           Logger
}

// New creates a Server. Pass repository.New(...) and log.New(...) from main.
func New(repo Repository, lgr Logger) *Server {
	return &Server{db: repo, lgr: lgr}
}

// ListenAndServe starts listening on the given port and blocks until the server stops.
func (s *Server) ListenAndServe(port int) error {
	const addr = "127.0.0.1"
	listener, err := net.Listen("tcp", net.JoinHostPort(addr, strconv.Itoa(port)))
	if err != nil {
		return fmt.Errorf("cannot listen on port %d: %w", port, err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterHabitsServer(grpcServer, s) // wires your Server struct to gRPC
	s.lgr.Logf("starting server on port %d\n", port)
	if err = grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("error while serving: %w", err)
	}
	return nil
}
