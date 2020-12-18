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

	passport := map[string]string{}
	validCount := 0

	for input.Scan() {
		line := input.Bytes()

		if len(strings.TrimSpace(string(line))) == 0 {
			if isValid(passport) {
				validCount++
			}
			passport = map[string]string{}
			continue
		}

		keyVals := strings.Split(string(line), " ")
		for _, keyVal := range keyVals {
			split := strings.Split(keyVal, ":")

			passport[split[0]] = split[1]
		}
	}

	fmt.Println(validCount)
}

func isValid(passport map[string]string) bool {
	byr, ok := passport["byr"]
	if !ok {
		return false
	}
	byrInt, err := strconv.Atoi(byr)
	if err != nil || 1920 > byrInt || byrInt > 2002 {
		fmt.Printf("birth year %s is not between 1920 and 2002\n", byr)
		return false
	}

	iyr, ok := passport["iyr"]
	if !ok {
		return false
	}
	iyrInt, err := strconv.Atoi(iyr)
	if err != nil || 2020 < iyrInt || iyrInt < 2010 {
		fmt.Printf("iyr %s is not between 2010 and 2020\n", iyr)
		return false
	}

	eyr, ok := passport["eyr"]
	if !ok {
		return false
	}
	eyrInt, err := strconv.Atoi(eyr)
	if err != nil || 2020 > eyrInt || eyrInt > 2030 {
		fmt.Printf("eyr %s is not between 2020 and 2030\n", eyr)
		return false
	}

	hgt, ok := passport["hgt"]
	if !ok {
		return false
	}
	if strings.HasSuffix(hgt, "cm") {
		hgtInt, err := strconv.Atoi(string([]byte(hgt)[0 : len(hgt)-2]))
		if err != nil || hgtInt < 150 || hgtInt > 193 {
			fmt.Printf("hgt %s is not between 150 and 193 cm\n", hgt)
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		hgtInt, err := strconv.Atoi(string([]byte(hgt)[0 : len(hgt)-2]))
		fmt.Printf("hgt %s is not between 59 and 76 in\n", hgt)
		if err != nil || hgtInt < 59 || hgtInt > 76 {
			return false
		}
	} else {
		return false
	}

	hcl, ok := passport["hcl"]
	if !ok {
		return false
	}
	if []byte(hcl)[0] != '#' {
		return false
	}
	for _, b := range []byte(hcl)[1:] {
		if (byte('9') >= b && byte('0') <= b) ||
			(byte('f') >= b && byte('a') <= b) {
			continue
		}
		fmt.Printf("'%s' is not an acceptable character in a hair color\n", string(b))
		return false
	}

	ecl, ok := passport["ecl"]
	if !ok {
		return false
	}
	foundEyeColor := false
	for _, eyeCol := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if eyeCol == ecl {
			foundEyeColor = true
			break
		}
	}
	if !foundEyeColor {
		fmt.Printf("eye color %s not approved\n", ecl)
		return false
	}

	pid, ok := passport["pid"]
	if !ok {
		return false
	}
	_, err = strconv.Atoi(pid)
	if err != nil || len(pid) != 9 {
		fmt.Printf("pid %s either not a number or not length 9", pid)
		return false
	}

	return true
}
