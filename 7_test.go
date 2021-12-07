package advent

import "testing"

func TestCrabDance(t *testing.T) {
	tables := []struct {
		input  []int
		answer int
	}{
		{input7example, 37},
		{input7, 344605},
	}

	for i, table := range tables {
		result := CrabDance(table.input)

		if result != table.answer {
			t.Errorf("Result was incorrect for line %d, expected: %d, got: %d.", i, table.answer, result)
		}
	}
}
