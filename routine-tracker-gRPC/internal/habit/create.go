package habit

import (
	"context"
	"fmt"
)

//go:generate minimock -i habitCreator -s "_mock.go" -o "mocks"
type habitCreator interface {
	Add(ctx context.Context, habit Habit) error
}

func Create(ctx context.Context, db habitCreator, h Habit) (Habit, error) {
	h, err := validateAndCompleteHabit(h)
	if err != nil {
		return Habit{}, err
	}
	err = db.Add(ctx, h)
	if err != nil {
		return Habit{}, fmt.Errorf("cannot save habit: %w", err)
	}
	return h, nil
}
