package habit

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// habitCreator is the only repository method Create needs.
// Small interfaces = easier to mock in tests.
//
//go:generate minimock -i habitCreator -s "_mock.go" -o "mocks"
type habitCreator interface {
	Add(ctx context.Context, habit Habit) error
}

// Create validates h, fills missing fields, saves it, and returns the saved habit.
func Create(ctx context.Context, db habitCreator, h Habit) (Habit, error) {
	h, err := validateAndFillDetails(h)
	if err != nil {
		return Habit{}, err
	}
	dbCtx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
	if err = db.Add(dbCtx, h); err != nil {
		return Habit{}, fmt.Errorf("cannot save habit: %w", err)
	}
	return h, nil
}

// validateAndFillDetails checks the input and fills server-side fields (ID, CreationTime).
func validateAndFillDetails(h Habit) (Habit, error) {
	h.Name = Name(strings.TrimSpace(string(h.Name)))
	if h.Name == "" {
		return Habit{}, InvalidInputError{field: "name", reason: "cannot be empty"}
	}
	if h.WeeklyFrequency == 0 {
		h.WeeklyFrequency = 1 // default to once a week
	}
	if h.ID == "" {
		h.ID = ID(uuid.NewString())
	}
	if h.CreationTime.IsZero() {
		h.CreationTime = time.Now()
	}
	return h, nil
}
