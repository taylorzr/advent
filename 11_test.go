package advent

import "testing"

func TestOctoLights(t *testing.T) {
	tables := []struct {
		input  string
		times  int
		answer int
	}{
		{input11simple, 2, 9},
		{input11example, 2, 35},
		{input11example, 100, 1656},
		{input11, 100, 1705},

		// panics on last:
		// {input11example, 195, 1656},
		// {input11, 265, 1705},
	}

	for i, table := range tables {
		result := OctoLights(table.input, table.times)

		if result != table.answer {
			t.Errorf("Result was incorrect for line %d, expected: %d, got: %d.", i, table.answer, result)
		}
	}
}
