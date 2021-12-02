package advent

import "fmt"

type Command struct {
	Direction string
	Distance  int
}

type Sub struct {
	X   int
	Y   int
	Aim int
}

func Move(commands []Command) int {
	sub := Sub{}

	for _, command := range commands {
		switch command.Direction {
		case "forward":
			sub.X += command.Distance
		case "down":
			sub.Y += command.Distance
		case "up":
			sub.Y -= command.Distance
		default:
			panic(fmt.Sprintf("invalid command: '%v'", command))
		}
	}

	// fmt.Printf("%v\n", sub)

	return sub.X * sub.Y
}

func MoveMore(commands []Command) int {
	sub := Sub{}

	/*
	   down X increases your aim by X units.
	   up X decreases your aim by X units.
	   forward X does two things:
	       It increases your horizontal sub by X units.
	       It increases your depth by your aim multiplied by X.
	*/

	for _, command := range commands {
		switch command.Direction {
		case "forward":
			sub.X += command.Distance
			sub.Y += command.Distance * sub.Aim
		case "down":
			sub.Aim += command.Distance
		case "up":
			sub.Aim -= command.Distance
		default:
			panic(fmt.Sprintf("invalid command: '%v'", command))
		}
	}

	// fmt.Printf("%v\n", sub)

	return sub.X * sub.Y
}
