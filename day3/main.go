package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	input := bufio.NewScanner(inputFile)

	var trees [][]bool

	for input.Scan() {
		line := input.Bytes()

		var treeLine []bool

		for _, b := range line {
			if b == '#' {
				treeLine = append(treeLine, true)
			} else {
				treeLine = append(treeLine, false)
			}
		}

		trees = append(trees, treeLine)
	}

	a := countTrees(trees, 1, 1)
	b := countTrees(trees, 3, 1)
	c := countTrees(trees, 5, 1)
	d := countTrees(trees, 7, 1)
	e := countTrees(trees, 1, 2)

	fmt.Println(a * b * c * d * e)
}

func countTrees(trees [][]bool, right, down int) int {
	mod := len(trees[0])
	x := 0
	y := 0
	treeCount := 0

	for y < len(trees) {
		if trees[y][x] {
			treeCount++
		}

		x = (x + right) % mod
		y += down
	}

	fmt.Println(treeCount)
	return treeCount
}
