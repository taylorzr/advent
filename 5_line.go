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
	// Vertical Line
	if line.Start.X == line.End.X {
		start, end := line.Start.Y, line.End.Y

		if line.Start.Y > line.End.Y {
			start, end = end, start
		}

		x := line.Start.X
		for i := start; i <= end; i++ {
			loc := Loc{X: x, Y: i}
			m.LocCounts[loc]++
		}
		return
	}

	// Horizontal Line
	if line.Start.Y == line.End.Y {
		start, end := line.Start.X, line.End.X

		if line.Start.X > line.End.X {
			start, end = end, start
		}

		y := line.Start.Y
		for i := start; i <= end; i++ {
			loc := Loc{X: i, Y: y}
			m.LocCounts[loc]++
		}

		return
	}

	// Diagonal Line
	// start, end := line.Start, line.End
	// if line.Start.X > line.End.X {
	// 	start, end = end, start
	// }

	// // fmt.Printf("%+v -> %+v\n", start, end)

	// x, y := start.X, start.Y
	// if start.Y > end.Y {
	// 	panic(fmt.Sprintf("bad assumption: %+v %+v\n", start, end))
	// }

	xd, yd := 1, 1
	if line.Start.X > line.End.X {
		xd = -1
	}
	if line.Start.Y > line.End.Y {
		yd = -1
	}

	x, y := line.Start.X, line.Start.Y
	for {
		loc := Loc{x, y}
		m.LocCounts[loc]++

		// fmt.Printf("%+v\n", loc)

		if x == line.End.X {
			break
		}

		x+=xd
		y+=yd
	}
}
