package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"adventofcode.com/AlexeyNeyman/day9"
)

func main() {
	fmt.Println(day9.Solve1(readAsset("day9_input.txt")))
}

func readAsset(name string) []byte {
	path, _ := filepath.Abs("./assets/" + name)

	contents, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return contents
}
