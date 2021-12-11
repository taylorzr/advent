package advent

import (
	"fmt"
	"strconv"
	"strings"
)

func OctoLights(input string, times int) int {
	octopi := NewOctopi(input)
	fmt.Println("Before any steps:")
	octopi.Print()

	flashes := 0
	for i := 0; i < times; i++ {
		for loc, octopus := range octopi.grid {
			if octopus.energy == 9 {
				octopi.Flash(loc)
			} else {
				// octopus.energy += 1
			}
		}

		// octopi.Print()

		// flash
		for _, octopus := range octopi.grid {
			if octopus.energy == 9 {
				flashes += 1
				octopus.energy = 0
				octopus.flashed = false
			} else {
				octopus.energy += 1
			}
		}

		fmt.Printf("After step %d:\n", i+1)
		octopi.Print()
	}

	return flashes
}

func NewOctopi(input string) *Octopi {
	octopi := Octopi{grid: map[Loc]*Octopus{}}

	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, digit := range line {
			if x == 0 && y == 0 {
				octopi.columns = len(line)
				octopi.rows = len(lines)
			}

			n, err := strconv.Atoi(string(digit))
			if err != nil {
				panic(err)
			}
			octopi.grid[Loc{x, y}] = &Octopus{energy: n}
		}
	}

	return &octopi
}

type Octopus struct {
	energy  int
	flashed bool
}

type Octopi struct {
	grid    map[Loc]*Octopus
	columns int
	rows    int
	flashes int
}

func (octopi *Octopi) Flash(loc Loc) {
	octopus := octopi.grid[loc]

	if octopus.flashed {
		return
	} else {
		octopi.flashes += 1
		octopus.flashed = true
	}

	neighborLocs := []Loc{
		// clockwise
		{loc.X + 0, loc.Y - 1}, // up
		{loc.X + 1, loc.Y - 1}, // up right
		{loc.X + 1, loc.Y + 0}, // right
		{loc.X + 1, loc.Y + 1}, // down right
		{loc.X + 0, loc.Y + 1}, // down
		{loc.X - 1, loc.Y + 1}, // down left
		{loc.X - 1, loc.Y + 0}, // left
		{loc.X - 1, loc.Y - 1}, // left up
	}

	for _, neighborLoc := range neighborLocs {
		neighbor, exists := octopi.grid[neighborLoc]

		if !exists {
			continue
		}

		if neighbor.energy != 9 {
			neighbor.energy += 1
		}

		if neighbor.energy == 9 {
			octopi.Flash(neighborLoc)
		}
	}
}

func (octopi *Octopi) Print() {
	for y := 0; y < octopi.rows; y++ {
		for x := 0; x < octopi.columns; x++ {
			octopus := octopi.grid[Loc{x, y}]
			fmt.Print(octopus.energy)
		}
		fmt.Print("\n")
	}
	fmt.Printf("--------\n")
}
