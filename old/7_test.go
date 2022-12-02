package advent

import "testing"

func TestCrabDance(t *testing.T) {
	tables := []struct {
		input  []int
		part2  bool
		answer int
	}{
		{input7example, false, 37},
		{input7, false, 344605},
		{input7example, true, 168},
		{input7, true, 93699985},
	}

	for i, table := range tables {
		result := CrabDance(table.input, table.part2)

		if result != table.answer {
			t.Errorf("Result was incorrect for line %d, expected: %d, got: %d.", i, table.answer, result)
		}
	}
}
