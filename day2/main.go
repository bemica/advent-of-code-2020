package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	input := bufio.NewScanner(inputFile)

	var validCount int

	for input.Scan() {
		line := input.Bytes()
		lineSplit := strings.Split(string(line), ":")
		rule := lineSplit[0]
		ruleSplit := strings.Split(rule, " ")
		minAndMax := strings.Split(ruleSplit[0], "-")
		char := ruleSplit[1]

		min, _ := strconv.Atoi(minAndMax[0])
		max, _ := strconv.Atoi(minAndMax[1])

		if isValidPass(strings.TrimSpace(lineSplit[1]), char, min, max) {
			fmt.Printf("%s is valid\n\n", string(line))
			validCount++
		} else {
			fmt.Printf("%s is invalid\n", string(line))
		}
	}

	fmt.Println(validCount)
}

func isValidPass(pass, char string, min, max int) bool {
	var indexes []int

	for i, c := range []byte(pass) {
		if c == []byte(char)[0] {
			indexes = append(indexes, i+1)
		}
	}

	foundOne := false

	for _, n := range indexes {
		if n == min {
			foundOne = true
		}

		if n == max {
			if foundOne {
				return false
			}

			return true
		}
	}

	return foundOne
}
