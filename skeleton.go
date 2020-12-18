package main

import (
	"bufio"
	"os"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	input := bufio.NewScanner(inputFile)

	for input.Scan() {
		line := input.Bytes()

	}
}
