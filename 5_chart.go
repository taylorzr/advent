package advent

import (
	"fmt"
	"strings"
)

func NewChart() *Chart {
	return &Chart{map[Loc]int{}}
}

type Chart struct {
	LocCounts map[Loc]int
}

func (chart *Chart) Draw(input string, skipDiagonal bool) {
	lineInputs := strings.Split(input, "\n")

	for _, lineInput := range lineInputs {
		line := NewLine(lineInput)
		if skipDiagonal && line.IsDiagonal() {
			continue
		}
		line.DrawOn(chart)
	}
}

func (chart *Chart) DangerCounts() int {
	fmt.Printf("%+v\n", chart.LocCounts)
	sum := 0
	for _, count := range chart.LocCounts {
		if count > 1 {
			sum++
		}
	}
	return sum
}
