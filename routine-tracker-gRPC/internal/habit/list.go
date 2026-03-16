package habit

import (
    "context"
    "fmt"
)

// habitLister is the only repository method ListHabits needs.
type habitLister interface {
    FindAll(ctx context.Context) ([]Habit, error)
}

// ListHabits returns all habits from the repository.
func ListHabits(ctx context.Context, db habitLister) ([]Habit, error) {
    habits, err := db.FindAll(ctx)
    if err != nil {
        return nil, fmt.Errorf("cannot list habits: %w", err)
    }
    return habits, nil
}