package day3

import (
	"adventofcode.com/AlexeyNeyman/util"
)

func Solve1(input []byte) int {
	var (
		powerConsumption int
		gammaRate        int
		epsilonRate      int
		numbers          []int = make([]int, 0)
		numberLength     int
	)

	for _, line := range util.ParseLines(string(input)) {
		if numberLength == 0 {
			numberLength = len(line)
		}

		number := util.ParseBinaryInt(line)
		numbers = append(numbers, number)
	}

	counter := newCounter(numberLength)

	counter.calcStats(numbers)

	for _, pos := range counter.positions() {
		mcb, lcb := counter.getBits(pos)
		shift := (numberLength - pos - 1)
		gammaRate = gammaRate | (mcb << shift)
		epsilonRate = epsilonRate | (lcb << shift)
	}

	powerConsumption = gammaRate * epsilonRate

	return powerConsumption
}

func Solve2(input []byte) int {
	var (
		lifeSupportRating     int
		oxygenGeneratorRating int
		co2ScrubberRating     int
		numbers               []int = make([]int, 0)
		numberLength          int
	)

	for _, line := range util.ParseLines(string(input)) {
		if numberLength == 0 {
			numberLength = len(line)
		}

		number := util.ParseBinaryInt(line)
		numbers = append(numbers, number)
	}

	oxygenGeneratorRating = findSpecialRatingValue(numbers, numberLength, true, 1)
	co2ScrubberRating = findSpecialRatingValue(numbers, numberLength, false, 0)
	lifeSupportRating = oxygenGeneratorRating * co2ScrubberRating

	return lifeSupportRating
}

func findSpecialRatingValue(numbers []int, numberLength int, mostCommonBit bool, preferredBit int) int {
	filteredNumbers := make([]int, len(numbers))
	copy(filteredNumbers, numbers)

	for pos := 0; pos < numberLength; pos++ {
		newFilteredNumbers := make([]int, 0)

		counter := newCounter(numberLength)
		counter.calcStats(filteredNumbers)

		targetBitValue := counter.getBit(pos, mostCommonBit, preferredBit)

		for _, num := range filteredNumbers {
			shift := (numberLength - pos - 1)
			bitValue := util.GetBitValue(num, shift)

			if bitValue == targetBitValue {
				newFilteredNumbers = append(newFilteredNumbers, num)
			}

		}

		filteredNumbers = newFilteredNumbers

		if len(filteredNumbers) == 1 {
			return filteredNumbers[0]
		}
	}

	return -1
}
