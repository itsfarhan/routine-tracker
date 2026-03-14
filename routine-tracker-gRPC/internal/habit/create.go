package habit

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/itsfarhan/routine-tracker/api"
)

// CreateHabit is the endpoint that registers a habit.
func (s *Server) CreateHabit(_ context.Context, request *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {

	s.lgr.Logf("CreateHabit request received: %s", request)
	return &api.CreateHabitResponse{Habit: &api.Habit{}}, nil
}

func validateAndCompleteHabit(h Habit) (Habit, error) {
	h.Name = Name(strings.TrimSpace(string(h.Name)))
	if h.Name == "" {
		return Habit{}, InvalidInputError{field: "name", reason: "cannot be empty"}
	}
	if h.WeeklyFrequency == 0 {
		h.WeeklyFrequency = 1 // default
	}
	if h.ID == "" {
		h.ID = ID(uuid.NewString())
	}
	if h.CreationTime.IsZero() {
		h.CreationTime = time.Now()
	}
	return h, nil
}
