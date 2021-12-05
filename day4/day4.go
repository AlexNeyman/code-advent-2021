package day4

import (
	"strings"

	"adventofcode.com/AlexeyNeyman/util"
)

func Solve1(input []byte) int {
	var score int

	numbersToDraw, boards := parseInput(string(input))

drawLoop:
	for _, n := range numbersToDraw {
		for _, board := range boards {
			board.mark(n)

			if board.isBingo() {
				score = board.score()
				break drawLoop
			}
		}
	}

	return score
}

func Solve2(input []byte) int {
	numbersToDraw, boards := parseInput(string(input))

	winningBoards := make([]*board, 0)

	for _, n := range numbersToDraw {
		stillPlayingBoards := make([]*board, 0)

		for _, board := range boards {
			board.mark(n)

			if board.isBingo() {
				winningBoards = append(winningBoards, board)
			} else {
				stillPlayingBoards = append(stillPlayingBoards, board)
			}
		}

		boards = stillPlayingBoards

		if len(boards) == 0 {
			break
		}
	}

	lastBoardToWin := winningBoards[len(winningBoards)-1]

	return lastBoardToWin.score()
}

func parseInput(input string) ([]int, []*board) {
	splittedInput := strings.Split(input, "\n\n")
	numbers := parseCommaSeparatedInts(splittedInput[0])
	boards := parseBoards(splittedInput[1:])

	return numbers, boards
}

func parseCommaSeparatedInts(str string) []int {
	nums := make([]int, 0)

	for _, s := range strings.Split(str, ",") {
		nums = append(nums, util.ParseInt(s))
	}

	return nums
}

func parseBoards(strs []string) []*board {
	boards := make([]*board, 0, len(strs))

	for _, str := range strs {
		boards = append(boards, parseBoard(str))
	}

	return boards
}
