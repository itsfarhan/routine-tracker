package habit_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/itsfarhan/routine-tracker/internal/habit"
	"github.com/itsfarhan/routine-tracker/internal/habit/mocks"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	h := habit.Habit{
		Name:            "swim",
		WeeklyFrequency: 2,
		CreationTime:    time.Now(),
		ID:              "123",
	}
	dbErr := fmt.Errorf("db unavailable")

	tests := map[string]struct {
		db          func(ctl *minimock.Controller) *mocks.HabitCreatorMock
		expectedErr error
	}{
		"nominal": {
			db: func(ctl *minimock.Controller) *mocks.HabitCreatorMock {
				db := mocks.NewHabitCreatorMock(ctl)
				db.AddMock.Expect(minimock.AnyContext, h).Return(nil)
				return db
			},
			expectedErr: nil,
		},
		"error case": {
			db: func(ctl *minimock.Controller) *mocks.HabitCreatorMock {
				db := mocks.NewHabitCreatorMock(ctl)
				db.AddMock.Expect(minimock.AnyContext, h).Return(dbErr)
				return db
			},
			expectedErr: dbErr,
		},
	}

	for name, tt := range tests {
		name, tt := name, tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := minimock.NewController(t)

			db := tt.db(ctrl)

			got, err := habit.Create(context.Background(), db, h)
			assert.ErrorIs(t, err, tt.expectedErr)
			if tt.expectedErr == nil {
				assert.Equal(t, h.Name, got.Name)
			}
		})
	}
}
