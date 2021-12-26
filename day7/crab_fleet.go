package day7

import "adventofcode.com/AlexeyNeyman/util"

type crabFleet struct {
	crabs []crab
	width int
}

func newCrabFleet(crabPositions []int, crabType submarine) *crabFleet {
	fleet := crabFleet{
		crabs: make([]crab, 0, len(crabPositions)),
	}

	for _, crabPosition := range crabPositions {
		fleet.crabs = append(fleet.crabs, newCrab(crabPosition, crabType))
		fleet.width = util.MaxInt(fleet.width, crabPosition)
	}

	return &fleet
}

func (cf *crabFleet) align() {
	possibleFuelConsumption := make([]int, 0, cf.width)

	for pos := 0; pos < cf.width; pos++ {
		possibleFuelConsumption =
			append(
				possibleFuelConsumption,
				cf.calculateAlignmentFuelConsumption(pos),
			)
	}

	bestAlignmentPosition := util.MinInt(possibleFuelConsumption...)

	for crabIndex := range cf.crabs {
		cf.moveCrab(crabIndex, bestAlignmentPosition)
	}
}

func (cf *crabFleet) calculateAlignmentFuelConsumption(position int) int {
	var fuelConsumption int

	for _, crab := range cf.crabs {
		fuelConsumption += crab.calculateMovementFuelConsumption(position)
	}

	return fuelConsumption
}

func (cf *crabFleet) moveCrab(crabIndex, newPosition int) {
	crab := cf.crabs[crabIndex]
	cf.crabs[crabIndex] = crab.moveTo(newPosition)
}
