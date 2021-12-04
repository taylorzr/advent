package advent

import (
	"fmt"
	"testing"
)

func TestBoardSimple(t *testing.T) {
	// tables := []struct {
	// 	input  []string
	// 	answer uint64
	// }{
	// 	{[]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}, 198},
	// 	{input3, 2954600},
	// }

	// for i, table := range tables {
	// 	result := Gamma(table.input)

	// 	if result != table.answer {
	// 		t.Errorf("Result was incorrect for line %d, got: %d, want: %d.", i, result, table.answer)
	// 	}
	// }

	picks := []int{57, 19, 40, 54, 64}
	input := `57 19 40 54 64
22 69 15 88  8
79 60 48 95 85
34 97 33  1 55
72 82 29 90 84`

	board := NewBoard(input)

	for i, pick := range picks {
		result := board.Mark(pick)

		if i < 4 && result == true {
			t.Errorf("Iteration %d, expected result to be false, but got true", i)
		}

		if i == 4 {
			if result != true {
				t.Errorf("Iteration %d, expected result to be false, but got true", i)
			}
			score := board.Score()
			if i == 4 && board.Score() != 73344 {
				t.Errorf("Iteration %d, expected score to be 73344, but got %d", i, score)
			}
		}
	}
}

func TestBoardExample(t *testing.T) {
	expectedWinRound := 11
	picks := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	input := `14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

	board := NewBoard(input)

	for i, pick := range picks {
		result := board.Mark(pick)

		if i < expectedWinRound && result == true {
			t.Errorf("Iteration %d, unexpected board winner", i)
		}

		if i == expectedWinRound {
			if result != true {
				fmt.Printf("%+v\n", board)
				t.Errorf("Iteration %d, expected board to have won", i)
			}
			score := board.Score()
			if i == expectedWinRound && board.Score() != 4512 {
				t.Errorf("Iteration %d, expected score to be 73344, but got %d", i, score)
			}

			break
		}
	}
}

func TestGameExample(t *testing.T) {
	game := NewGame(input4example)
	result := game.Play()

	if result.Score() != 4512 {
		t.Errorf("Expected winner score of 4512, but got %d", result.Score())
	}
}

func TestGameFull(t *testing.T) {
	game := NewGame(input4)

	for _, b := range game.Boards {
		fmt.Printf("%s\n\n", b.Input)
	}

	result := game.Play()

	if result.Score() != 11774 {
		t.Errorf("Expected winner score of 11774, but got %d", result.Score())
	}
}

func TestGameFullAll(t *testing.T) {
	game := NewGame(input4)

	for _, b := range game.Boards {
		fmt.Printf("%s\n\n", b.Input)
	}

	result := game.PlayAll()

	if result.Score() != 4495 {
		t.Errorf("Expected winner score of 4495, but got %d", result.Score())
	}

}
