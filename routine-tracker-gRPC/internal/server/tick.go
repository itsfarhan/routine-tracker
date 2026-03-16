package server

import (
	"context"
	"errors"
	"time"

	"github.com/itsfarhan/routine-tracker/api"
	"github.com/itsfarhan/routine-tracker/internal/habit"
	"github.com/itsfarhan/routine-tracker/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) TickHabit(ctx context.Context, req *api.TickHabitRequest) (*api.TickHabitResponse, error) {
	err := habit.Tick(ctx, s.db, s.db, habit.ID(req.HabitId), time.Now())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "habit %q not found", req.HabitId)
		}
		return nil, status.Errorf(codes.Internal, "cannot tick habit: %s", err)
	}
	return &api.TickHabitResponse{}, nil
}
