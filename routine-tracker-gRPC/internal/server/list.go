package server

import (
	"context"

	"github.com/itsfarhan/routine-tracker/api"
	"github.com/itsfarhan/routine-tracker/internal/habit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListHabits(ctx context.Context, _ *api.ListHabitsRequest) (*api.ListHabitsResponse, error) {
	habits, err := habit.ListHabits(ctx, s.db)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot list habits: %s", err)
	}
	apiHabits := make([]*api.Habit, len(habits))
	for i, h := range habits {
		apiHabits[i] = &api.Habit{
			Id:              string(h.ID),
			Name:            string(h.Name),
			WeeklyFrequency: int32(h.WeeklyFrequency),
		}
	}
	return &api.ListHabitsResponse{Habits: apiHabits}, nil
}
