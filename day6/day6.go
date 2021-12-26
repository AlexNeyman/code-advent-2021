package day6

import "adventofcode.com/AlexeyNeyman/util"

func Solve1(input []byte) int {
	return predictFishPopulation(input, 80)
}

func Solve2(input []byte) int {
	return predictFishPopulation(input, 256)
}

func predictFishPopulation(input []byte, afterDays int) int {
	initialFish := util.ParseCommaSeparatedInts(string(input))

	simulation := newFishSimulation(initialFish)

	simulation.passDays(afterDays)

	fishCountAtTheEnd := simulation.countFish()

	return fishCountAtTheEnd
}
