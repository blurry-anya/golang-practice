package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	currentDate := time.Now()
	parsedDate, _ := time.Parse("January 2, 2006 15:04:05", date)
	isPassed := parsedDate.Compare(currentDate)
	return isPassed < 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := parsedDate.Hour()
	if hour >= 12 && hour < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedDate, _ := time.Parse("1/2/2006 15:04:05", date)

	return fmt.Sprintf("You have an appointment on %v, %v %d, %d, at %d:%d.", parsedDate.Weekday(), parsedDate.Month(), parsedDate.Day(), parsedDate.Year(), parsedDate.Hour(), parsedDate.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
