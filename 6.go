package advent

func allBirths(daysUntil []int, remainingDays int) int {
	sum := 0
	for _, days := range daysUntil {
		sum += births(days, remainingDays)
	}
	return sum
}

func births(daysUntil, remainingDays int) int {
	remainingDays -= daysUntil

	if remainingDays <= 0 {
		return 1
	}

	return births(7, remainingDays) + births(9, remainingDays)
}
