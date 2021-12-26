package day8

type displayEntry struct {
	signalPatterns []string
	outputValues   []string
}

func (e displayEntry) countUniqueLengthOutputValues() int {
	var result int

	for _, outputValue := range e.outputValues {
		switch len(outputValue) {
		case 2, 3, 4, 7:
			result += 1
		}
	}

	return result
}
