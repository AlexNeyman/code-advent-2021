package util

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func NewLinesScanner(str string) *bufio.Scanner {
	reader := strings.NewReader(str)
	scanner := bufio.NewScanner(reader)

	return scanner
}

func ParseLines(str string) []string {
	lines := make([]string, 0)

	for scanner := NewLinesScanner(str); scanner.Scan(); {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func ParseInt(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		log.Fatal(err)
	}

	return num
}
func ParseBinaryInt(str string) int {
	num, err := strconv.ParseInt(str, 2, 0)

	if err != nil {
		log.Fatal(err)
	}

	return int(num)
}

func ParseIntLines(str string) []int {
	nums := make([]int, 0)

	for scanner := NewLinesScanner(str); scanner.Scan(); {
		line := scanner.Text()
		num := ParseInt(line)
		nums = append(nums, num)
	}

	return nums
}

func ParseSeparatedInts(str string, sep string) []int {
	nums := make([]int, 0)

	for _, s := range strings.Split(str, sep) {
		nums = append(nums, ParseInt(s))
	}

	return nums
}

func ParseCommaSeparatedInts(str string) []int {
	return ParseSeparatedInts(str, ",")
}

func ParseSeparatedIntLines(str string, sep string) [][]int {
	intLines := make([][]int, 0)

	for scanner := NewLinesScanner(str); scanner.Scan(); {
		line := scanner.Text()
		intLine := ParseSeparatedInts(line, sep)
		intLines = append(intLines, intLine)
	}

	return intLines
}
