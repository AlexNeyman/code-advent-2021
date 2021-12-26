package day7

import "adventofcode.com/AlexeyNeyman/util"

type submarine interface {
	calculateMovementFuelConsumption(int, int) int
}

type linearFuelConsumptionSubmarine struct{}

type increasingFuelConsumptionSubmarine struct{}

type crab struct {
	position  int
	submarine submarine
}

func newCrab(position int, submarine submarine) crab {
	return crab{
		position:  position,
		submarine: submarine,
	}
}

func (c crab) moveTo(newPosition int) crab {
	c.position = newPosition

	return c
}

func (c crab) calculateMovementFuelConsumption(newPosition int) int {
	return c.submarine.calculateMovementFuelConsumption(c.position, newPosition)
}

func (s linearFuelConsumptionSubmarine) calculateMovementFuelConsumption(from, to int) int {
	return util.AbsInt(from - to)
}

func (s increasingFuelConsumptionSubmarine) calculateMovementFuelConsumption(from, to int) int {
	var fuelConsumption int

	distance := util.AbsInt(from - to)

	for i := 0; i < distance; i++ {
		fuelConsumption += 1 + i
	}

	return fuelConsumption
}
