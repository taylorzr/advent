package advent

import (
	"fmt"
	"strconv"
)

func Gamma(inputs []string) uint64 {
	gamma := make([]rune, len(inputs[0]))
	epsilon := make([]rune, len(inputs[0]))

OUT:
	for i := 0; i < len(inputs[0]); i++ {
		gamma[i] = '0'
		epsilon[i] = '1'
		oneCount := 0

		for _, input := range inputs {
			if input[i] == '1' {
				oneCount++
			}

			if oneCount > len(inputs)/2 {
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

func LifeSupport(inputs []string) uint64 {
	filtered := inputs

	for i := 0; i < len(inputs[0]); i++ {
		oneCount := 0
		ones := []string{}
		zeros := []string{}

		for _, input := range filtered {
			if input[i] == '1' {
				oneCount++
				ones = append(ones, input)
			} else {
				zeros = append(zeros, input)
			}
		}

		// fmt.Printf("i: %d l: %v 1-count: %d length: %d\n", i, filtered, oneCount, len(filtered))

		if oneCount >= len(filtered)-oneCount {
			filtered = ones
		} else {
			filtered = zeros
		}

		if len(filtered) == 1 {
			break
		}
	}

	oxygen := filtered[0]

	fmt.Printf("oxygen: %s\n", oxygen)

	o, err := strconv.ParseUint(oxygen, 2, 32)
	if err != nil {
		panic(err)
	}

	// e, err := strconv.ParseUint(string(epsilon), 2, 32)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Printf("o: %d\n", o)

	return o
}
