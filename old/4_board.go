package advent

import (
	"strconv"
	"strings"
)

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

func (board *Board) Mark(n int) bool {
	item, exists := board.Items[n]
	if !exists {
		return false
	}

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
