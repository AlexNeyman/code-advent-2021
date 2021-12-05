package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"adventofcode.com/AlexeyNeyman/day4"
)

func main() {
	fmt.Println(day4.Solve2(readAsset("day4_input.txt")))
}

func readAsset(name string) []byte {
	path, _ := filepath.Abs("./assets/" + name)

	contents, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return contents
}
