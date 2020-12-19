package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
)

func main() {
	inputFile, _ := os.Open("inputTest.txt")
	input := bufio.NewScanner(inputFile)

	jolts := []int{0}
	myJoltage := 3

	for input.Scan() {
		line := input.Bytes()
		jolt, _ := strconv.Atoi(string(line))

		if jolt+3 > myJoltage {
			myJoltage = jolt + 3
		}

		jolts = append(jolts, jolt)
	}
	jolts = append(jolts, myJoltage)
	sort.Ints(jolts)

	//diffCount := map[int]int{}
	//prevJolt := jolts[0]
	//for _, jolt := range jolts[1:] {
	//	_, ok := diffCount[jolt-prevJolt]
	//	if !ok {
	//		diffCount[jolt-prevJolt] = 0
	//	}
	//
	//	diffCount[jolt-prevJolt]++
	//
	//	prevJolt = jolt
	//}

	fmt.Println(countArrangements(jolts, map[string]*big.Int{}))
}

func countArrangements(jolts []int, cache map[string]*big.Int) *big.Int {
	subArrangements := big.NewInt(0)
	one := big.NewInt(1)

	for i := range jolts[1 : len(jolts)-1] {
		if jolts[i+2]-jolts[i] <= 3 {
			newJolts := removeAndCopy(jolts, i+1)
			//fmt.Printf("cutting %d out of %v (%dth index) to form %v\n", jolts[i+1], jolts, i+1, newJolts)
			if _, ok := cache[getString(jolts)]; ok {
				//subArrangements.Add(val, subArrangements)
				continue
			}

			calculated := countArrangements(
				newJolts,
				cache,
			)
			subArrangements.Add(
				calculated,
				subArrangements,
			)

			subArrangements.Add(one, subArrangements)
		}
	}

	cache[getString(jolts)] = subArrangements

	fmt.Printf("found %v sub arrangement of %v\n", subArrangements, jolts)
	return subArrangements
}

func getString(jolts []int) string {
	return fmt.Sprintf("%v", jolts)
}

func removeAndCopy(jolts []int, index int) []int {
	var arr []int

	for i := range jolts {
		if i == index {
			continue
		}

		arr = append(arr, jolts[i])
	}

	return arr
}
