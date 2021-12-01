package advent

import "testing"

func TestCountIncreasing(t *testing.T) {
	tables := []struct {
		input  []int
		answer int
	}{
		{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 7},
		{input1, 1215},
	}

	for i, table := range tables {
		count := CountIncreasing(table.input)

		if count != table.answer {
			t.Errorf("Count was incorrect for line %d, got: %d, want: %d.", i, count, table.answer)
		}
	}
}

func TestCountIncreasingRange(t *testing.T) {
	tables := []struct {
		input  []int
		answer int
	}{
		{[]int{607, 618, 618, 617, 647, 716, 769, 792}, 5},
		{input1, 1150},
	}

	for i, table := range tables {
		count := CountIncreasingWindow(table.input)

		if count != table.answer {
			t.Errorf("Count was incorrect for line %d, got: %d, want: %d.", i, count, table.answer)
		}
	}
}
