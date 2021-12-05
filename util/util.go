package util

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func GetBitValue(num int, shift int) int {
	return (num & (1 << shift)) >> shift
}

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

func SumInts(nums []int) int {
	res := 0

	for _, num := range nums {
		res += num
	}

	return res
}
