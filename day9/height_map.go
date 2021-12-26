package day9

import (
	"adventofcode.com/AlexeyNeyman/util"
)

type heightMap [][]int

func (hm heightMap) getLowPoints() []int {
	lowPoints := make([]int, 0)

	for i, row := range hm {
		for j, point := range row {
			adjancedPoints := hm.getAdjancedPoints(i, j)

			if util.MinInt(append(adjancedPoints, point)...) == point {
				lowPoints = append(lowPoints, point)
			}
		}
	}

	return lowPoints
}

func (hm heightMap) getAdjancedPoints(targetI, targetJ int) []int {
	adjancedPoints := make([]int, 0)

	for i := targetI - 1; i <= targetI+1; i++ {
		if i < 0 || i > len(hm)-1 {
			continue
		}

		row := hm[i]

		for j := targetJ - 1; j <= targetJ+1; j++ {
			if j < 0 || j > len(row)-1 {
				continue
			}

			if i == targetI && j == targetJ {
				continue
			}

			point := row[j]

			adjancedPoints = append(adjancedPoints, point)
		}
	}

	return adjancedPoints
}

func (hm heightMap) getRiskLevels() []int {
	riskLevels := make([]int, 0)

	for _, lowPoint := range hm.getLowPoints() {
		riskLevels = append(riskLevels, getRiskLevel(lowPoint))
	}

	return riskLevels
}

func (hm heightMap) sumRiskLevels() int {
	return util.SumInts(hm.getRiskLevels())
}

func getRiskLevel(point int) int {
	return 1 + point
}
