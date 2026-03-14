// Package habit holds the domain types and business logic for habits.
package habit

import (
	"time"
)

type ID string
type Name string
type WeeklyFrequency uint

type Habit struct {
	ID              ID
	Name            Name
	WeeklyFrequency WeeklyFrequency
	CreationTime    time.Time // internal field, not in the API
}
