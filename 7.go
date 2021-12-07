package advent

import (
	"fmt"
	"math"
	"sort"
)

func CrabDance(input []int, part2 bool) int {
	positions := make([]int, len(input))
	copy(positions, input)
	sort.Slice(positions, func(i, j int) bool { return positions[i] < positions[j] })
	median := positions[len(positions)/2]

	sum := 0
	for _, n := range positions {
		sum += n
	}
	avg := int(math.Round(float64(sum) / float64(len(positions))))

	fmt.Printf("median: %d avg: %d sum: %d\n", median, avg, sum)

	direction := 1
	if avg < median {
		direction = -1
	}

	n := median
	minN := median
	minFuel := math.MaxFloat32
	for {
		fuel := 0.0
		// n := positions[i]

		for _, position := range positions {
			if !part2 {
				fuel += math.Abs(float64(position) - float64(n))
			} else {
				diff := math.Abs(float64(position) - float64(n))
				fuel += (diff*diff + diff) / 2
			}
		}
		fmt.Printf("n: %d fuel %d\n", n, int(fuel))
		if fuel <= minFuel {
			minN = n
			minFuel = fuel
		} else {
			fmt.Printf("minN: %d minFuel: %d\n", minN, int(minFuel))
			break
		}
		// i += direction
		n += direction

	}

	return int(minFuel)
}
