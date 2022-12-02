package advent

import (
	// "fmt"
	"sort"
	"strings"
)

var opener = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

func ScoreSyntaxErrors(input string) int {
	scoring := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	score := 0

	for _, line := range strings.Split(input, "\n") {
		// fmt.Printf("line: %s\n", line)

		stack := []rune{}

		for _, char := range line {
			// fmt.Printf("stack: %v\n", stack)
			switch char {
			case '(', '[', '{', '<':
				stack = append(stack, char)
			case ')', ']', '}', '>':
				if stack[len(stack)-1] != opener[char] {
					score += scoring[char]
					break
				}
				stack = stack[0 : len(stack)-1]
			}
		}
	}

	return score
}

func ScoreSyntaxCompletion(input string) int {
	scoring := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	scores := []int{}

LINES:
	for _, line := range strings.Split(input, "\n") {
		// fmt.Printf("line: %s\n", line)

		stack := []rune{}

		for _, char := range line {
			// fmt.Printf("stack: %v\n", stack)
			switch char {
			case '(', '[', '{', '<':
				stack = append(stack, char)
			case ')', ']', '}', '>':
				if stack[len(stack)-1] != opener[char] {
					continue LINES
				}
				stack = stack[0 : len(stack)-1]
			}
		}

		// We should now have only opening chars
		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			opener := stack[i]
			score *= 5
			score += scoring[opener]
		}

		scores = append(scores, score)
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}
