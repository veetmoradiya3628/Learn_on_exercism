package booking

import (
    "time"
	"fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	appTime, err := time.Parse("January 2, 2006 15:04:05", date)
    if err != nil {
        panic(err)
    }
    now := time.Now()
    difference := appTime.Sub(now)
    return difference < 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
    return t.Hour() >= 12 && t.Hour() < 18;
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	   layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
    return fmt.Sprintf("You have an appointment on %s, at %s.",
		t.Format("Monday, January 2, 2006"),
		t.Format("15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	 year := time.Now().Year()
	return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}
