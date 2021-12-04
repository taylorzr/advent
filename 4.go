package advent

import (
	"fmt"
	"strconv"
	"strings"
)

type Item struct {
	X    int
	Y    int
	Seen bool
}

type Board struct {
	Input      string
	ColumnHits []int
	RowHits    []int
	Items      map[int]*Item
	Winner     int
	Won        bool
}

func NewBoard(input string) *Board {
	board := Board{
		Input: input,
		Items: map[int]*Item{},
	}

	lines := strings.Split(input, "\n")
	board.ColumnHits = make([]int, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)

		if i == 0 {
			board.RowHits = make([]int, len(parts))
		}

		for j, part := range parts {
			n, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}

			board.Items[n] = &Item{i, j, false}
		}

	}

	return &board
}

func (board *Board) Mark(n int) bool {
	_, exists := board.Items[n]
	if !exists {
		return false
	}

	item := board.Items[n]
	item.Seen = true

	board.RowHits[item.X]++
	board.ColumnHits[item.Y]++

	if board.RowHits[item.X] == len(board.ColumnHits) {
		// fmt.Printf("Winner: row %d\n", item.X)
		board.Winner, board.Won = n, true
		return true
	}

	if board.ColumnHits[item.Y] == len(board.RowHits) {
		// fmt.Printf("Winner: column %d\n", item.Y)
		board.Winner, board.Won = n, true
		return true
	}

	return false
}

func (board *Board) Score() int {
	sum := 0

	for n, item := range board.Items {
		if !item.Seen {
			sum += n
		}
	}

	return sum * board.Winner
}

type Game struct {
	Boards  []*Board
	Winners []int
	Picks   []int
}

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