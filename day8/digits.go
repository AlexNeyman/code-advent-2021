package day8

import (
	"log"
	"strings"
)

// 0:      1:      2:      3:      4:
// aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
// ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
// gggg    ....    gggg    gggg    ....

//  5:      6:      7:      8:      9:
// aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
// dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
// gggg    gggg    ....    gggg    gggg

type signalPattern struct {
	a bool
	b bool
	c bool
	d bool
	e bool
	f bool
	g bool
}

var (
	zero  = newSignalPattern("abcefg")
	one   = newSignalPattern("cf")
	two   = newSignalPattern("acdeg")
	three = newSignalPattern("acdfg")
	four  = newSignalPattern("bcdf")
	five  = newSignalPattern("abdfg")
	six   = newSignalPattern("abdefg")
	seven = newSignalPattern("acf")
	eight = newSignalPattern("abcdefg")
	nine  = newSignalPattern("abcdfg")
)

var digits = map[signalPattern]int{
	zero:  0,
	one:   1,
	two:   2,
	three: 3,
	four:  4,
	five:  5,
	six:   6,
	seven: 7,
	eight: 8,
	nine:  9,
}

var signalDigits = splitSignalPatternString("abcdefg")

func newSignalPattern(signals string) signalPattern {
	var sp signalPattern

	for _, signal := range splitSignalPatternString(signals) {
		switch signal {
		case "a":
			sp.a = true
		case "b":
			sp.b = true
		case "c":
			sp.c = true
		case "d":
			sp.d = true
		case "e":
			sp.e = true
		case "f":
			sp.f = true
		case "g":
			sp.g = true
		default:
			log.Panic("unexpected signal", signal)
		}
	}

	return sp
}

func splitSignalPatternString(str string) []string {
	return strings.Split(str, "")
}
