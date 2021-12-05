package advent

import (
	// "fmt"
	"strconv"
	"strings"
)

func NewLine(input string) Line {
	line := Line{Start: Loc{}, End: Loc{}}
	input = strings.ReplaceAll(input, " -> ", ",")

	for i, part := range strings.Split(input, ",") {
		n, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		switch i {
		case 0:
			line.Start.X = n
		case 1:
			line.Start.Y = n
		case 2:
			line.End.X = n
		case 3:
			line.End.Y = n
		}
	}

	return line
}

type Loc struct {
	X int
	Y int
}

type Line struct {
	Start Loc
	End   Loc
}

func (line *Line) IsDiagonal() bool {
	if line.Start.X == line.End.X || line.Start.Y == line.End.Y {
		return false
	}
	return true
}

func (line *Line) DrawOn(m *Chart) {
	xd, yd := 1, 1

	if line.Start.X == line.End.X {
		xd = 0
	}
	if line.Start.X > line.End.X {
		xd = -1
	}

	if line.Start.Y == line.End.Y {
		yd = 0
	}
	if line.Start.Y > line.End.Y {
		yd = -1
	}

	x, y := line.Start.X, line.Start.Y
	for {
		loc := Loc{x, y}
		m.LocCounts[loc]++

		if x == line.End.X && y == line.End.Y {
			break
		}

		x+=xd
		y+=yd
	}
}
