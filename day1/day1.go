package day1

import (
	"adventofcode.com/AlexeyNeyman/util"
)

func Solve1(input []byte) int {
	var (
		prevValue      int
		increasedCount int
	)

	for _, value := range util.ParseIntLines(string(input)) {
		if prevValue > 0 && prevValue < value {
			increasedCount += 1
		}

		prevValue = value
	}

	return increasedCount
}

func Solve2(input []byte) int {
	var (
		prevWindowSum  int
		increasedCount int
	)

	values := util.ParseIntLines(string(input))
	valueWindows := groupWindows(values, 3)

	for _, window := range valueWindows {
		windowSum := util.SumInts(window)

		if prevWindowSum > 0 && prevWindowSum < windowSum {
			increasedCount += 1
		}

		prevWindowSum = windowSum
	}

	return increasedCount
}

func groupWindows(xs []int, windowSize int) [][]int {
	windows := make([][]int, 0)

	if len(xs) <= windowSize {
		return append(windows, xs)
	}

	lastWindow := xs[0:windowSize]
	windows = append(windows, lastWindow)

	for i := windowSize; i < len(xs); i++ {
		x := xs[i]
		window := append(lastWindow[1:], x)
		windows = append(windows, window)
		lastWindow = window
	}

	return windows
}
