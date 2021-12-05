package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"adventofcode.com/AlexeyNeyman/day5"
)

func main() {
	fmt.Println(day5.Solve2(readAsset("day5_input.txt")))
}

func readAsset(name string) []byte {
	path, _ := filepath.Abs("./assets/" + name)

	contents, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return contents
}
