package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirds int
	for i := 0; i < len(birdsPerDay); i++ {
		totalBirds += birdsPerDay[i]
	}
	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var totalBirds int
	days := 0 + (week * 6)
	if week <= 1 {
		for i := 0; i <= days; i++ {
			totalBirds += birdsPerDay[i]
		}
	} else if week > 1 {
		for i := days + 1; i >= days-5; i-- {
			totalBirds += birdsPerDay[i]
		}
	}

	return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	var totalBirds = birdsPerDay
	for i := 0; i < len(birdsPerDay); i++ {
		if i == 0 {
			totalBirds[i] += 1
		} else if (i+1)%2 == 0 {
			totalBirds[i] += 0
		} else {
			totalBirds[i] += 1
		}
	}
	return totalBirds
}
