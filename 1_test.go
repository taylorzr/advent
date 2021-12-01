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

func TestCountIncreasingRangeOf3(t *testing.T) {
	tables := []struct {
		input  []int
		answer int
	}{
		{[]int{607, 618, 618, 617, 647, 716, 769, 792}, 5},
		{input1, 1150},
	}

	for i, table := range tables {
		count := CountIncreasingWindowOf3(table.input)

		if count != table.answer {
			t.Errorf("Count was incorrect for line %d, got: %d, want: %d.", i, count, table.answer)
		}
	}
}

func TestCountIncreasingWindow(t *testing.T) {
	tables := []struct {
		input  []int
		window int
		answer int
	}{
		{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 1, 7},
		{input1, 1, 1215},
		{[]int{607, 618, 618, 617, 647, 716, 769, 792}, 3, 5},
		{input1, 3, 1150},
	}

	for i, table := range tables {
		count := CountIncreasingWindow(table.input, table.window)

		if count != table.answer {
			t.Errorf("Count was incorrect for line %d, got: %d, want: %d.", i, count, table.answer)
		}
	}
}
