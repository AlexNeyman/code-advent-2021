package day8

import (
	"strings"

	"adventofcode.com/AlexeyNeyman/util"
)

func Solve1(input []byte) int {
	var result int

	lines := util.ParseLines(string(input))

	for _, line := range lines {
		displayEntry := parseDisplayEntry(line)

		result += displayEntry.countUniqueLengthOutputValues()
	}

	return result
}

// func Solve2(input []byte) int {
// 	var result int

// 	lines := util.ParseLines(string(input))

// 	for _, line := range lines {
// 		displayEntry := parseDisplayEntry(line)

// 		p := prmt.New(prmt.StringSlice(signalDigits))

// 		displayEntry.use
// 	}

// 	return result
// }

func parseDisplayEntry(str string) displayEntry {
	splittedStr := strings.Split(str, " | ")
	signalPatterns := strings.Split(splittedStr[0], " ")
	outputValues := strings.Split(splittedStr[1], " ")

	return displayEntry{
		signalPatterns: signalPatterns,
		outputValues:   outputValues,
	}
}
