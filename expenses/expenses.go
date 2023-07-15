package expenses

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	filtered := []Record{}
	for _, v := range in {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	count := 0.0
	filtered := Filter(in, ByDaysPeriod(p))
	for _, v := range filtered {
		count += v.Amount
	}
	return count
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	inCategory := Filter(in, ByCategory(c))
	if len(inCategory) == 0 {
		errMessage := fmt.Sprintf("unknown category %v", c)
		return 0, errors.New(errMessage)
	}
	inPeriod := Filter(inCategory, ByDaysPeriod(p))
	if len(inPeriod) == 0 {
		return 0, nil
	}
	count := 0.0
	for _, v := range inPeriod {
		count += v.Amount
	}
	return count, nil
}
