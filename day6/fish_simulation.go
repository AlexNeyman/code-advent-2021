package day6

type fishSimulation struct {
	population fishPopulation
}

type fishPopulation map[fish]int

type fish int

func newFishSimulation(initialfish []int) *fishSimulation {
	Simulation := fishSimulation{
		population: make(fishPopulation, 0),
	}

	Simulation.addFish(initialfish)

	return &Simulation
}

func (fs *fishSimulation) addOneFish(oneFish int) {
	fs.population[fish(oneFish)] += 1
}

func (fs *fishSimulation) addFish(fish []int) {
	for _, oneFish := range fish {
		fs.addOneFish(oneFish)
	}
}

func (fs *fishSimulation) passDays(days int) {
	for i := 0; i < days; i++ {
		fs.passDay()
	}
}

func (fs *fishSimulation) passDay() {
	fs.population = fs.population.afterDay()
}

func (fs *fishSimulation) countFish() int {
	var fishTotalCount int

	for _, fishCount := range fs.population {
		fishTotalCount += fishCount
	}

	return fishTotalCount
}

func (fp fishPopulation) afterDay() fishPopulation {
	newPopulation := make(fishPopulation)

	for fish, fishCount := range fp {
		for _, newFish := range fish.afterDay() {
			newPopulation[newFish] += fishCount
		}
	}

	return newPopulation
}

func (f fish) afterDay() []fish {
	if f == 0 {
		return []fish{6, 8}
	}

	return []fish{f - 1}
}
