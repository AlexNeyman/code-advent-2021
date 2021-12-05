package day4

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"

	"adventofcode.com/AlexeyNeyman/util"
)

type board struct {
	numbers      []*boardNumber
	drawnNumbers []int
}

type boardNumber struct {
	n        int
	isMarked bool
}

const (
	boardRowLength    = 5
	boardRowCount     = 5
	boardColumnLength = boardRowCount
	boardColumnCount  = boardRowLength
	boardNumberCount  = boardRowLength * boardColumnLength
)

func newBoard(numbers []int) *board {
	boardNumbers := make([]*boardNumber, 0, boardNumberCount)
	drawnNumbers := make([]int, 0)

	for _, n := range numbers {
		boardNumbers = append(boardNumbers, &boardNumber{n: n, isMarked: false})
	}

	return &board{
		numbers:      boardNumbers,
		drawnNumbers: drawnNumbers,
	}
}

func parseBoard(str string) *board {
	numbers := make([]int, 0)

	for _, numberRow := range strings.Split(str, "\n") {
		for _, n := range parseNumberRow(numberRow) {
			numbers = append(numbers, n)
		}
	}

	return newBoard(numbers)
}

func parseNumberRow(str string) []int {
	numberRow := make([]int, 0)
	re := regexp.MustCompile("\\d+")

	for _, s := range re.FindAllString(str, -1) {
		numberRow = append(numberRow, util.ParseInt(s))
	}

	return numberRow
}

func (b *board) mark(n int) {
	b.drawnNumbers = append(b.drawnNumbers, n)

	for _, boardNumber := range b.numbers {
		if boardNumber.n == n {
			boardNumber.isMarked = true
			return
		}
	}
}

func (b *board) isBingo() bool {
	return b.anyRowIsMarked() || b.anyColumnIsMarked()
}

func (b *board) anyRowIsMarked() bool {
	return anySequenceIsMarked(b.rows())
}

func (b *board) anyColumnIsMarked() bool {
	return anySequenceIsMarked(b.columns())
}

func (b *board) rows() [][]*boardNumber {
	rows := make([][]*boardNumber, 0, boardRowCount)

	for i := 0; i < boardRowCount; i++ {
		from := i * boardRowLength
		to := from + boardRowLength
		rows = append(rows, b.numbers[from:to])
	}

	return rows
}

func (b *board) columns() [][]*boardNumber {
	columns := make([][]*boardNumber, 0, boardRowLength)

	for i := 0; i < boardColumnCount; i++ {
		column := make([]*boardNumber, 0, boardColumnLength)

		for j := 0; j < boardColumnLength; j++ {
			column = append(column, b.numbers[i+j*boardRowLength])
		}

		columns = append(columns, column)
	}

	return columns
}

func (b *board) score() int {
	return b.getLastDrawnNumber() * b.sumAllUnmarkedNumbers()
}

func (b *board) getLastDrawnNumber() int {
	drawnNumbersCount := len(b.drawnNumbers)

	if drawnNumbersCount == 0 {
		log.Panic("no drawn numbers")
	}

	return b.drawnNumbers[drawnNumbersCount-1]
}

func (b *board) sumAllUnmarkedNumbers() int {
	var sum int

	for _, number := range b.numbers {
		if !number.isMarked {
			sum += number.n
		}
	}

	return sum
}

func (b *board) render() string {
	var buf bytes.Buffer

	for _, row := range b.rows() {
		for _, number := range row {
			fmt.Fprintf(&buf, "%2d ", number.n)
		}
		buf.WriteString("\n")
	}

	return buf.String()
}

func anySequenceIsMarked(seqs [][]*boardNumber) bool {
	for _, seq := range seqs {
		allNumbersAreMarked := true

		for _, number := range seq {
			if !number.isMarked {
				allNumbersAreMarked = false
				break
			}
		}

		if allNumbersAreMarked {
			return true
		}
	}

	return false
}
