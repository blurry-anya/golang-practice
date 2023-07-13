package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
	for _, v := range birdsPerDay {
		count += v
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	count := 0

	neededWeek := []int{}
	currentWeek := 0
	for i := 0; i < len(birdsPerDay); i += 7 {
		currentWeek++
		if currentWeek == week {
			if len(birdsPerDay) == 7 {
				neededWeek = append(neededWeek, birdsPerDay...)
				break
			} else {
				neededWeek = append(neededWeek, birdsPerDay[i:i+8]...)
				break
			}
		}
	}

	for _, v := range neededWeek {
		count += v
	}
	return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	birdsPerDay[0]++
	for i := 2; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}

	return birdsPerDay
}
