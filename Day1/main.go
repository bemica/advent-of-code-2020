package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	input := bufio.NewScanner(inputFile)

	var vals []int

	for input.Scan() {
		line := input.Bytes()

		n, err := strconv.Atoi(string(line))
		if err != nil {
			panic(err)
		}

		vals = append(vals, n)
	}

	for i, n := range vals {
		for j := i + 1; j < len(vals); j++ {
			m := vals[j]

			for k := j + 1; k < len(vals); k++ {
				if n+m+vals[k] == 2020 {
					fmt.Println(n * m * vals[k])
				}
			}
		}
	}
}
