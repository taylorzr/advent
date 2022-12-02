package advent

import "testing"

func TestSyntaxErrors(t *testing.T) {
	tables := []struct {
		input  string
		answer int
	}{
		{input10example, 26397},
		{input10, 339537},
	}

	for i, table := range tables {
		result := ScoreSyntaxErrors(table.input)

		if result != table.answer {
			t.Errorf("Result was incorrect for line %d, expected: %d, got: %d.", i, table.answer, result)
		}
	}
}

func TestSyntaxCompletion(t *testing.T) {
	tables := []struct {
		input  string
		answer int
	}{
		{input10example, 288957},
		{input10, 2412013412},
	}

	for i, table := range tables {
		result := ScoreSyntaxCompletion(table.input)

		if result != table.answer {
			t.Errorf("Result was incorrect for line %d, expected: %d, got: %d.", i, table.answer, result)
		}
	}
}
