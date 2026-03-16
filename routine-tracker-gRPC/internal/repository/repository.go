package repository

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/itsfarhan/routine-tracker/internal/habit"
	"github.com/itsfarhan/routine-tracker/internal/isoweek"
)

// ErrNotFound is returned when a habit ID doesn't exist.
const ErrNotFound = repositoryError("habit not found")

type repositoryError string

func (e repositoryError) Error() string { return string(e) }

// Logger matches the same interface used everywhere else.
type Logger interface {
	Logf(format string, args ...any)
}

// HabitRepository is the in-memory store for habits.
type HabitRepository struct {
	mu     sync.Mutex
	habits map[habit.ID]habit.Habit
	ticks  map[habit.ID]map[isoweek.ISO8601][]time.Time
	lgr    Logger
}

// New returns an empty, ready-to-use HabitRepository.
func New(lgr Logger) *HabitRepository {
	return &HabitRepository{
		habits: make(map[habit.ID]habit.Habit),
		ticks:  make(map[habit.ID]map[isoweek.ISO8601][]time.Time),
		lgr:    lgr,
	}
}

// Add stores a new habit. Overwrites if the ID already exists.
func (r *HabitRepository) Add(_ context.Context, h habit.Habit) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.habits[h.ID] = h
	r.lgr.Logf("stored habit %s", h.ID)
	return nil
}

// FindAll returns all habits sorted by creation time (deterministic order).
func (r *HabitRepository) FindAll(_ context.Context) ([]habit.Habit, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	habits := make([]habit.Habit, 0, len(r.habits))
	for _, h := range r.habits {
		habits = append(habits, h)
	}
	sort.Slice(habits, func(i, j int) bool {
		return habits[i].CreationTime.Before(habits[j].CreationTime)
	})
	return habits, nil
}

// Find returns a single habit by ID.
func (r *HabitRepository) Find(_ context.Context, id habit.ID) (habit.Habit, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	h, ok := r.habits[id]
	if !ok {
		return habit.Habit{}, fmt.Errorf("%w: %s", ErrNotFound, id)
	}
	return h, nil
}

// AddTick records a tick for a habit at the given time, grouped by ISO week.
func (r *HabitRepository) AddTick(_ context.Context, id habit.ID, t time.Time) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.habits[id]; !ok {
		return fmt.Errorf("%w: %s", ErrNotFound, id)
	}
	if r.ticks[id] == nil {
		r.ticks[id] = make(map[isoweek.ISO8601][]time.Time)
	}
	year, week := t.ISOWeek()
	key := isoweek.ISO8601{Year: year, Week: week}
	r.ticks[id][key] = append(r.ticks[id][key], t)
	return nil
}
