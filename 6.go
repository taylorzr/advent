package advent

func allBirths(daysUntil []int, remainingDays int) int {

	daysToCounts := map[int]int{}

	for _, day := range daysUntil {
		daysToCounts[day] += 1
	}

	sum := 0
	for days, count := range daysToCounts {
		sum += births(days, remainingDays) * count
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

func iBirths(input []int, remainingDays int) int {
	fishes := make([]int, len(input))
	copy(fishes, input)

	for i := remainingDays; i > 0; i-- {
		newFishes := []int{}

		for i, fish := range fishes {
			if fish == 0 {
				newFishes = append(newFishes, 8)
				fishes[i] = 6
			} else {
				fishes[i] -= 1
			}
		}

		// fmt.Printf("%v %v\n", fishes, newFishes)
		fishes = append(fishes, newFishes...)
	}

	return len(fishes)
}
