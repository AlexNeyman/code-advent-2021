package day6

import "adventofcode.com/AlexeyNeyman/util"

func Solve1(input []byte) int {
	return simulateFishPopulation(input, 80)
}

func Solve2(input []byte) int {
	return simulateFishPopulation(input, 256)
}

func simulateFishPopulation(input []byte, simulationLengthDays int) int {
	initialFish := util.ParseCommaSeparatedInts(string(input))

	simulation := newFishSimulation(initialFish)

	simulation.passDays(simulationLengthDays)

	fishCountAtTheEnd := simulation.countFish()

	return fishCountAtTheEnd
}
