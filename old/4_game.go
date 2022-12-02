package advent

import (
	"fmt"
	"strconv"
	"strings"
)

func NewGame(input string) *Game {
	game := Game{}

	sections := strings.Split(input, "\n\n")

	picksInput := sections[0]

	for _, pickInput := range strings.Split(picksInput, ",") {
		n, err := strconv.Atoi(pickInput)
		if err != nil {
			panic(err)
		}
		game.Picks = append(game.Picks, n)
	}

	boardInputs := sections[1:]
	game.Boards = make([]*Board, len(boardInputs))
	game.Winners = make([]int, len(boardInputs))

	for i, boardInput := range boardInputs {
		game.Boards[i] = NewBoard(boardInput)
	}

	return &game
}

type Game struct {
	Boards  []*Board
	Winners []int
	Picks   []int
}

func (game *Game) Play() *Board {
	for i, pick := range game.Picks {
		fmt.Printf("Round %d...%d\n", i+1, pick)
		for j, board := range game.Boards {
			win := board.Mark(pick)

			if win {
				fmt.Printf("Winner: board %d\n", j)
				fmt.Printf("Board:\n%s\n", board.Input)
				fmt.Printf("Board: %+v\n", board)
				return board
			}
		}
	}

	return nil
}

func (game *Game) PlayAll() *Board {
	winners := 0

	for i, pick := range game.Picks {
		fmt.Printf("Round %d...%d\n", i+1, pick)

		for j, board := range game.Boards {
			if board.Won {
				continue
			}

			win := board.Mark(pick)

			if win {
				winners++
				if winners == len(game.Boards) {
					fmt.Printf("Loser: board %d\n", j)
					return game.Boards[j]
				}

				fmt.Printf("Winner: board %d\n", j)
				// last := len(game.Boards) - 1
				// game.Boards[j] = game.Boards[last]
				// game.Boards[last] = nil
				// game.Boards = game.Boards[:last]
			}
		}
	}

	return nil
}
