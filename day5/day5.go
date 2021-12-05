package day5

import (
	"strings"
)

func Solve1(input []byte) int {
	diagram := newDiagram()
	lines := parseLines(string(input))

	for _, line := range lines {
		if line.isHorizontal() || line.isVertical() {
			diagram.addLine(line)
		}
	}

	dangerPointsCount := diagram.countPointsGreaterThan(1)

	return dangerPointsCount
}

func Solve2(input []byte) int {
	diagram := newDiagram()
	lines := parseLines(string(input))

	for _, line := range lines {
		diagram.addLine(line)
	}

	dangerPointsCount := diagram.countPointsGreaterThan(1)

	return dangerPointsCount
}

func parseLines(input string) []line {
	inputLines := strings.Split(input, "\n")
	lines := make([]line, 0, len(inputLines))

	for _, inputLine := range inputLines {
		lines = append(lines, parseLine(inputLine))
	}

	return lines
}
