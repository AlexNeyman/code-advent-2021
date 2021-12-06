package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"adventofcode.com/AlexeyNeyman/day6"
)

func main() {
	fmt.Println(day6.Solve2(readAsset("day6_input.txt")))
}

func readAsset(name string) []byte {
	path, _ := filepath.Abs("./assets/" + name)

	contents, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return contents
}
