package server

import (
	"context"
	"errors"

	"github.com/itsfarhan/routine-tracker/api"
	"github.com/itsfarhan/routine-tracker/internal/habit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateHabit(ctx context.Context, req *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	var freq uint
	if req.WeeklyFrequency != nil {
		freq = uint(*req.WeeklyFrequency)
	}
	h := habit.Habit{
		Name:            habit.Name(req.Name),
		WeeklyFrequency: habit.WeeklyFrequency(freq),
	}
	created, err := habit.Create(ctx, s.db, h)
	if err != nil {
		var invalidErr habit.InvalidInputError
		if errors.As(err, &invalidErr) {
			return nil, status.Error(codes.InvalidArgument, invalidErr.Error())
		}
		return nil, status.Errorf(codes.Internal, "cannot save habit: %s", err)
	}
	return &api.CreateHabitResponse{
		Habit: &api.Habit{
			Id:              string(created.ID),
			Name:            string(created.Name),
			WeeklyFrequency: int32(created.WeeklyFrequency),
		},
	}, nil
}
