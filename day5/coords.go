package day5

import "adventofcode.com/AlexeyNeyman/util"

type coords struct {
	x int
	y int
}

func parseCoords(input string) coords {
	xAndY := util.ParseCommaSeparatedInts(input)
	x := xAndY[0]
	y := xAndY[1]

	return coords{x: x, y: y}
}

func (c coords) stepTo(target coords) coords {
	xDir := decideDirection(c.x, target.x)
	yDir := decideDirection(c.y, target.y)

	return coords{
		x: c.x + xDir,
		y: c.y + yDir,
	}
}

func decideDirection(from, to int) int {
	if from == to {
		return 0
	} else if from < to {
		return 1
	} else {
		return -1
	}
}
