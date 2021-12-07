package advent

import (
	"fmt"
	"math"
	"sort"
)

func CrabDance(input []int) int {
	positions := make([]int, len(input))
	copy(positions, input)
	sort.Slice(positions, func(i, j int) bool { return positions[i] < positions[j] })
	medianIndex := len(positions) / 2
	median := positions[medianIndex]

	fmt.Printf("%v\n", positions)

	sum := 0
	for _, n := range positions {
		sum += n
	}
	avg := sum / len(positions)

	fmt.Printf("median: %d avg: %d sum: %d\n", median, avg, sum)

	direction := 1
	if avg < median {
		direction = -1
	}

	i := medianIndex
	minFuel := math.MaxFloat32
	for {
		fuel := 0.0
		n := positions[i]

		for _, position := range positions {
			fuel += math.Abs(float64(position) - float64(n))
		}
		fmt.Printf("i: %d n: %d fuel %d\n", i, n, int(fuel))
		if fuel <= minFuel {
			minFuel = fuel
		} else {
			break
		}
		i += direction

	}

	return int(minFuel)
}
