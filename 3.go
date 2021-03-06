package advent

import (
	"fmt"
	"strconv"
)

func PowerConsumption(inputs []string) uint64 {
	gamma := make([]rune, len(inputs[0]))
	epsilon := make([]rune, len(inputs[0]))

OUT:
	for i := range inputs[0] {
		gamma[i] = '0'
		epsilon[i] = '1'
		oneCount := 0

		for _, input := range inputs {
			if input[i] == '1' {
				oneCount++
			}

			if oneCount >= len(inputs)-oneCount {
				gamma[i] = '1'
				epsilon[i] = '0'
				continue OUT
			}

		}
	}

	fmt.Printf("gamma: %s epsilon: %s\n", string(gamma), string(epsilon))

	g, err := strconv.ParseUint(string(gamma), 2, 32)
	if err != nil {
		panic(err)
	}

	e, err := strconv.ParseUint(string(epsilon), 2, 32)
	if err != nil {
		panic(err)
	}

	fmt.Printf("g: %d e: %d\n", g, e)

	return g * e
}

type tracker struct {
	OneCount    int
	ZeroCount   int
	OneIndexes  map[int]bool
	ZeroIndexes map[int]bool
}

func NewTracker() *tracker {
	return &tracker{OneIndexes: map[int]bool{}, ZeroIndexes: map[int]bool{}}
}

func LifeSupport(inputs []string) uint64 {
	oxygenIndexes := map[int]bool{}
	co2Indexes := map[int]bool{}

	for i := range inputs[0] {
		oxygenTracking := NewTracker()
		co2Tracking := NewTracker()

		for j, input := range inputs {
			// We keep track of indexes we want to include in the results
			// so only count if the index is being tracked
			// and always count the first iteration because the initial map is empty
			if i == 0 || oxygenIndexes[j] && len(oxygenIndexes) != 1 {
				if input[i] == '1' {
					oxygenTracking.OneCount++
					oxygenTracking.OneIndexes[j] = true
				} else {
					oxygenTracking.ZeroCount++
					oxygenTracking.ZeroIndexes[j] = true
				}
			}

			if i == 0 || co2Indexes[j] && len(co2Indexes) != 1 {
				if input[i] == '1' {
					co2Tracking.OneCount++
					co2Tracking.OneIndexes[j] = true
				} else {
					co2Tracking.ZeroCount++
					co2Tracking.ZeroIndexes[j] = true
				}
			}
		}

		if i == 0 || len(oxygenIndexes) > 1 {
			if oxygenTracking.OneCount >= oxygenTracking.ZeroCount {
				oxygenIndexes = oxygenTracking.OneIndexes
			} else {
				oxygenIndexes = oxygenTracking.ZeroIndexes
			}
		}

		if i == 0 || len(co2Indexes) > 1 {
			if co2Tracking.OneCount < co2Tracking.ZeroCount {
				co2Indexes = co2Tracking.OneIndexes
			} else {
				co2Indexes = co2Tracking.ZeroIndexes
			}
		}

		if len(oxygenIndexes) == 1 && len(co2Indexes) == 1 {
			break
		}
	}

	oxygen := ""
	for i := range oxygenIndexes {
		oxygen = inputs[i]
	}

	co2 := ""
	for i := range co2Indexes {
		co2 = inputs[i]
	}

	o, err := strconv.ParseUint(oxygen, 2, 32)
	if err != nil {
		panic(err)
	}

	c, err := strconv.ParseUint(co2, 2, 32)
	if err != nil {
		panic(err)
	}

	fmt.Printf("oxygen: %s -> %d co2: %s -> %d\n", oxygen, o, co2, c)

	return o * c
}

func LifeSupportSeparated(inputs []string) uint64 {
	oxygen := findCommon(inputs, func(a, b int) bool { return a >= b })
	co2 := findCommon(inputs, func(a, b int) bool { return a < b })

	o, err := strconv.ParseUint(oxygen, 2, 32)
	if err != nil {
		panic(err)
	}

	c, err := strconv.ParseUint(co2, 2, 32)
	if err != nil {
		panic(err)
	}

	fmt.Printf("oxygen: %s -> %d co2: %s -> %d\n", oxygen, o, co2, c)

	return c * o
}

func findCommon(inputs []string, compare func(int, int) bool) string {
	filtered := inputs
	i := 0

	for len(filtered) > 1 {
		onesCount := 0
		zerosCount := 0
		ones := []string{}
		zeros := []string{}

		for _, input := range filtered {
			if input[i] == '1' {
				onesCount++
				ones = append(ones, input)
			} else {
				zerosCount++
				zeros = append(zeros, input)
			}
		}

		if compare(onesCount, zerosCount) {
			filtered = ones
		} else {
			filtered = zeros
		}

		i++
	}

	return filtered[0]
}
