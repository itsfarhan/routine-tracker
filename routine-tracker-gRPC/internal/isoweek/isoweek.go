// Package isoweek provides a type to identify a week by ISO 8601 year and week number.
package isoweek

// ISO8601 identifies a specific week in a specific year.
// Use time.Time.ISOWeek() to get the year and week values.
type ISO8601 struct {
	Year int
	Week int
}
