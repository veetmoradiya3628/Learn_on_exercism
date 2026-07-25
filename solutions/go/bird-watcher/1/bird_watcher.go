package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	ans := 0
    for idx := 0; idx < len(birdsPerDay); idx++ {
        ans += birdsPerDay[idx]
    }
    return ans
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    startIdx := (week - 1) * 7
	birdsInWeek := birdsPerDay[startIdx:startIdx+7]
    ans := 0
    for _, val := range birdsInWeek {
        ans += val
    }
    return ans
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for idx := 0; idx < len(birdsPerDay); idx++ {
        if idx % 2 == 0 {
            birdsPerDay[idx]++
        }
    }
    return birdsPerDay
}
