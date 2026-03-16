package habit

import (
	"context"
	"fmt"
	"time"
)

// habitFinder checks a habit exists before ticking it.
type habitFinder interface {
	Find(ctx context.Context, id ID) (Habit, error)
}

// tickAdder stores the tick in the repository.
type tickAdder interface {
	AddTick(ctx context.Context, id ID, t time.Time) error
}

// Tick records that a habit was done at time t.
// It first checks the habit exists, then stores the tick.
func Tick(ctx context.Context, habitDB habitFinder, tickDB tickAdder, id ID, t time.Time) error {
	if _, err := habitDB.Find(ctx, id); err != nil {
		return fmt.Errorf("cannot find habit %q: %w", id, err)
	}
	if err := tickDB.AddTick(ctx, id, t); err != nil {
		return fmt.Errorf("cannot tick habit %q: %w", id, err)
	}
	return nil
}
