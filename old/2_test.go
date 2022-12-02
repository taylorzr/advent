package advent

import "testing"

func TestMove(t *testing.T) {
	example := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	tables := []struct {
		input  []Command
		answer int
	}{
		{input2Parse(example), 150},
		{input2Parse(input2), 1815044},
	}

	for i, table := range tables {
		result := Move(table.input)

		if result != table.answer {
			t.Errorf("Position was incorrect for line %d, got: %d, want: %d.", i, result, table.answer)
		}
	}
}

func TestMoveMore(t *testing.T) {
	example := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	tables := []struct {
		input  []Command
		answer int
	}{
		{input2Parse(example), 900},
		{input2Parse(input2), 1739283308},
	}

	for i, table := range tables {
		result := MoveMore(table.input)

		if result != table.answer {
			t.Errorf("Position was incorrect for line %d, got: %d, want: %d.", i, result, table.answer)
		}
	}
}
