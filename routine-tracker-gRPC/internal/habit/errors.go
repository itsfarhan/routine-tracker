package habit

import "fmt"

// Invalidinputerror is returned when user i/p data is invalid
type InvalidInputError struct {
	field  string
	reason string
}

// Error implements error
func (e InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input in field %s: %s", e.field, e.reason)
}
