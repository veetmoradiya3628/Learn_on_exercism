package bafflingbirthdays

import (
	"math/rand"
	"time"
)

// SharedBirthday checks if any two dates in the slice share the same month and day.
func SharedBirthday(dates []time.Time) bool {
	seen := make(map[string]bool)
	
	for _, date := range dates {
		key := date.Format("01-02") 
		if seen[key] {
			return true
		}
		seen[key] = true
	}
	return false
}

// RandomBirthdates generates a slice of random dates of the specified size.
func RandomBirthdates(size int) []time.Time {
	dates := make([]time.Time, size)
	startOfYear := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	
	for i := 0; i < size; i++ {
		// FIXED: Changed back to 365 to perfectly match the test suite's math baseline
		randomDay := rand.Intn(365)
		dates[i] = startOfYear.AddDate(0, 0, randomDay)
	}
	return dates
}

// EstimatedProbability runs a simulation to approximate the probability 
// of a shared birthday within a given group size.
func EstimatedProbability(size int) float64 {
	const totalTrials = 50000
	sharedCount := 0

	for i := 0; i < totalTrials; i++ {
		if SharedBirthday(RandomBirthdates(size)) {
			sharedCount++
		}
	}

	// FIXED: Multiplied by 100.0 to return a percentage instead of a decimal fraction
	return (float64(sharedCount) / float64(totalTrials)) * 100.0
}
