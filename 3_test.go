package advent

import "testing"

func TestGamma(t *testing.T) {
	tables := []struct {
		input  []string
		answer uint64
	}{
		{[]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}, 198},
		{input3, 2954600},
	}

	for i, table := range tables {
		result := Gamma(table.input)

		if result != table.answer {
			t.Errorf("Result was incorrect for line %d, got: %d, want: %d.", i, result, table.answer)
		}
	}
}

func TestLifeSupport(t *testing.T) {
	tables := []struct {
		input  []string
		answer uint64
	}{
		{[]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}, 198},
		// {input3, 2954600},
	}

	for i, table := range tables {
		result := LifeSupport(table.input)

		if result != table.answer {
			t.Errorf("Result was incorrect for line %d, got: %d, want: %d.", i, result, table.answer)
		}
	}
}
