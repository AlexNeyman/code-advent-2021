package day5

import (
	"bytes"
	"fmt"

	"adventofcode.com/AlexeyNeyman/util"
)

type diagram map[coords]point

func newDiagram() diagram {
	return make(diagram)
}

func (d diagram) addLine(line line) {
	currentCoord := line.from

	for {
		d[currentCoord] += 1

		if currentCoord == line.to {
			break
		}

		currentCoord = currentCoord.stepTo(line.to)
	}
}

func (d diagram) rows() [][]point {
	rows := make([][]point, 0)
	maxX := d.sizeX()
	maxY := d.sizeY()

	for y := 0; y <= maxY; y++ {
		row := make([]point, 0)

		for x := 0; x <= maxX; x++ {
			pointCoords := coords{x: x, y: y}
			point := d[pointCoords]
			row = append(row, point)
		}

		rows = append(rows, row)
	}

	return rows
}

func (d diagram) sizeX() int {
	var maxX int

	for coords := range d {
		maxX = util.Max(coords.x, maxX)
	}

	return maxX
}

func (d diagram) sizeY() int {
	var maxY int

	for coords := range d {
		maxY = util.Max(coords.y, maxY)
	}

	return maxY
}

func (d diagram) countPointsGreaterThan(val point) int {
	var count int

	for _, row := range d.rows() {
		for _, point := range row {
			if point > val {
				count += 1
			}
		}
	}

	return count
}

func (d diagram) String() string {
	var buf bytes.Buffer

	for _, row := range d.rows() {
		for _, point := range row {
			fmt.Fprint(&buf, point, " ")
		}
		fmt.Fprint(&buf, "\n")
	}

	return buf.String()
}
