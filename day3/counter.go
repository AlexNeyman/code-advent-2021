package day3

import (
	"sort"

	"adventofcode.com/AlexeyNeyman/util"
)

type counter struct {
	stats        map[int]map[int]int
	numberLength int
}

func newCounter(numberLength int) *counter {
	return &counter{
		stats:        make(map[int]map[int]int),
		numberLength: numberLength,
	}
}

func (c *counter) calcStats(numbers []int) {
	for _, num := range numbers {

		for pos := 0; pos < c.numberLength; pos++ {
			bitValue := util.GetBitValue(num, c.numberLength-pos-1)

			// fmt.Printf("num: %d=%b, pos: %d, bitValue: %d\n", num, num, pos, bitValue)

			c.inc(pos, bitValue)
		}
	}
}

func (c *counter) inc(pos int, value int) {
	if c.stats[pos] == nil {
		c.stats[pos] = make(map[int]int)
	}

	c.stats[pos][value] += 1
}

func (c *counter) positions() []int {
	positions := make([]int, 0)

	for position := range c.stats {
		positions = append(positions, position)
	}

	sort.Ints(positions)

	return positions
}

func (c *counter) getBits(pos int) (int, int) {
	bits := c.stats[pos]

	if bits[1] >= bits[0] {
		return 1, 0
	} else {
		return 0, 1
	}
}

func (c *counter) getBit(pos int, mostCommonBit bool, preferredBit int) int {
	bits := c.stats[pos]
	zeros := bits[0]
	ones := bits[1]

	if zeros == ones {
		return preferredBit
	}

	if mostCommonBit {
		if zeros > ones {
			return 0
		} else {
			return 1
		}
	} else {
		if zeros < ones {
			return 0
		} else {
			return 1
		}
	}
}
