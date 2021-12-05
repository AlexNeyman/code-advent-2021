package day2

import (
	"strings"

	"adventofcode.com/AlexeyNeyman/util"
)

type direction string

const (
	directionUp      direction = "up"
	directionDown              = "down"
	directionForward           = "forward"
)

func Solve1(input []byte) int {
	var (
		position int
		depth    int
	)

	for _, line := range util.ParseLines(string(input)) {
		direction, distance := parseCourseDirective(line)

		switch direction {
		case directionUp:
			depth -= distance
		case directionDown:
			depth += distance
		case directionForward:
			position += distance
		}
	}

	return position * depth
}

func Solve2(input []byte) int {
	var (
		position int
		depth    int
		aim      int
	)

	for _, line := range util.ParseLines(string(input)) {
		direction, distance := parseCourseDirective(line)

		switch direction {
		case directionUp:
			aim -= distance
		case directionDown:
			aim += distance
		case directionForward:
			position += distance
			depth += aim * distance
		}
	}

	return position * depth
}

func parseCourseDirective(directive string) (direction, int) {
	splitted := strings.Split(directive, " ")

	direction := direction(splitted[0])
	distance := util.ParseInt(splitted[1])

	return direction, distance
}
