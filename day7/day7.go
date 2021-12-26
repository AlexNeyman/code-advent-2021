package day7

import "adventofcode.com/AlexeyNeyman/util"

func Solve1(input []byte) int {
	return calculateBestCrabsAlignmentPosition(
		input,
		linearFuelConsumptionSubmarine{},
	)
}

func Solve2(input []byte) int {
	return calculateBestCrabsAlignmentPosition(
		input,
		increasingFuelConsumptionSubmarine{},
	)
}

func calculateBestCrabsAlignmentPosition(
	input []byte,
	submarine submarine,
) int {
	initialCrabPositions :=
		util.ParseCommaSeparatedInts(string(input))

	crabFleet := newCrabFleet(initialCrabPositions, submarine)

	crabFleet.align()

	firstCrab := crabFleet.crabs[0]

	return firstCrab.position
}
