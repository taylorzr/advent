package advent

import "testing"

func TestSquidReproduction(t *testing.T) {
	tables := []struct {
		daysUntil     []int
		remainingDays int
		answer        int
	}{
		{input6example, 18, 26},
		{input6example, 80, 5934},
		{input6, 80, 390011},
		// These work but take forever, prolly cause go doesn't have tail call optimization
		// Or I'm missing a better way to solve this
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

/*
$ go test -run TestSquidReproduction -timeout 99999s
--- FAIL: TestSquidReproduction (5158.28s)
    6_test.go:22: Result was incorrect for line 3, expected: 0, got: 1746710169834.
FAIL
exit status 1
FAIL    github.com/taylorzr/advent      5158.651s
*/
