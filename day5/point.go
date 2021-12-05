package day5

import "strconv"

type point int

func (p point) String() string {
	if p == 0 {
		return "."
	}

	return strconv.Itoa(int(p))
}
