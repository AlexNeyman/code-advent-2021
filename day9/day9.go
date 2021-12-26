package day9

import "adventofcode.com/AlexeyNeyman/util"

func Solve1(input []byte) int {
	heightMap := parseHeightMap(input)

	return heightMap.sumRiskLevels()
}

func parseHeightMap(input []byte) heightMap {
	return heightMap(util.ParseSeparatedIntLines(string(input), ""))
}
