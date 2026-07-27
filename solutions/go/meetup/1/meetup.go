package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)
func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	var startDay int

	switch wSched {
	case First:
		startDay = 1
	case Second:
		startDay = 8
	case Third:
		startDay = 15
	case Fourth:
		startDay = 22
	case Teenth:
		startDay = 13
	case Last:
		// Start searching from 7 days before the last day of the month
		startDay = daysInMonth(month, year) - 6
	}

	// Find the matching wDay starting from our calculated baseline day
	for d := startDay; ; d++ {
		t := time.Date(year, month, d, 0, 0, 0, 0, time.UTC)
		if t.Weekday() == wDay {
			return t.Day()
		}
	}
}
func daysInMonth(m time.Month, y int) int {
	// Adding 1 to the month and using day 0 gives the last day of the current month
	return time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

