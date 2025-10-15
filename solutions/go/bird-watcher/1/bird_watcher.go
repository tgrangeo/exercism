package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    ret := 0
	for i := 0; i < len(birdsPerDay); i++ {
        ret += birdsPerDay[i]
    }
	return ret
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    start := (week - 1) * 7
    end := start + 7
	ret := 0
	for start < end {
        ret += birdsPerDay[start]
        start++
    }
	return ret
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, value := range birdsPerDay{
        if (i % 2 == 0){
        	birdsPerDay[i] = value + 1
        }
    }
	return birdsPerDay
}
