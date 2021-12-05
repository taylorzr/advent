package advent

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLine(t *testing.T) {
	input := "9,4 -> 3,4"
	line := NewLine(input)

	fmt.Printf("%+v\n", line)

	if !reflect.DeepEqual(line.Start, Loc{9, 4}) {
		panic("wrong")
	}
	if !reflect.DeepEqual(line.End, Loc{3, 4}) {
		panic("wrong")
	}

	chart := NewChart()
	chart.Draw(input, true)

	fmt.Printf("%+v\n", chart.LocCounts)
}

func TestLineDiagonal(t *testing.T) {
	input := "6,4 -> 2,0"

	// line := NewLine(input)
	chart := NewChart()
	chart.Draw(input, false)
	fmt.Printf("%+v\n", chart.LocCounts)

	// 6,4 -> 2,0
	// 0,0 -> 8,8
}

func TestChartExample(t *testing.T) {
	tables := []struct {
		input        string
		skipDiagonal bool
		answer       int
	}{
		{input5example, true, 5},
		{input5, true, 5092},
		{input5example, false, 12},
		{input5, false, 20484},
	}

	for i, table := range tables {
		chart := NewChart()
		chart.Draw(table.input, table.skipDiagonal)
		result := chart.DangerCounts()

		if result != table.answer {
			t.Errorf("Result was incorrect for line %d, got: %d, want: %d.", i, result, table.answer)
		}
	}
}
