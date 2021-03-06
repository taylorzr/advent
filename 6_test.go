package advent

import "testing"

func TestSquidReproductionRecursive(t *testing.T) {
	tables := []struct {
		daysUntil     []int
		remainingDays int
		answer        int
	}{
		{input6example, 18, 26},
		{input6example, 80, 5934},
		{input6, 80, 390011},
		// These are a bit slow ~ 1 min each
		// {input6example, 256, 26984457539},
		// {input6, 256, 1746710169834},
	}

	for i, table := range tables {
		result := allBirths(table.daysUntil, table.remainingDays)

		if result != table.answer {
			t.Errorf("Result was incorrect for line %d, expected: %d, got: %d.", i, table.answer, result)
		}
	}
}

func TestSquidReproductionIterative(t *testing.T) {
	tables := []struct {
		daysUntil     []int
		remainingDays int
		answer        int
	}{
		{input6example, 18, 26},
		{input6example, 80, 5934},
		{input6, 80, 390011},
		// {input6example, 256, 26984457539},
		// {input6, 256, 1746710169834},
	}

	for i, table := range tables {
		result := iBirths(table.daysUntil, table.remainingDays)

		if result != table.answer {
			t.Errorf("Result was incorrect for line %d, expected: %d, got: %d.", i, table.answer, result)
		}
	}
}
