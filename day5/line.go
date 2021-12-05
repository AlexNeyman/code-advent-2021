package day5

import "strings"

type line struct {
	from coords
	to   coords
}

func parseLine(input string) line {
	fromAndTo := strings.Split(input, " -> ")
	from := parseCoords(fromAndTo[0])
	to := parseCoords(fromAndTo[1])

	return line{from: from, to: to}
}

func (l line) isHorizontal() bool {
	return l.from.x == l.to.x
}

func (l line) isVertical() bool {
	return l.from.y == l.to.y
}
