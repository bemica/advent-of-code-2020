package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	input := bufio.NewScanner(inputFile)

	var seats []int64

	for input.Scan() {
		line := input.Bytes()
		rowStr := convertToBinary(line[:7], 'F', 'B')
		row, err := strconv.ParseInt(rowStr, 2, 64)
		if err != nil {
			panic(fmt.Errorf("could not convert %s to binary int", rowStr))
		}

		colStr := convertToBinary(line[7:], 'L', 'R')
		col, err := strconv.ParseInt(colStr, 2, 64)
		if err != nil {
			panic(fmt.Errorf("could not convert %s to binary int", rowStr))
		}

		seatId := row*8 + col
		seats = append(seats, seatId)
	}

	for _, seat := range seats {
		for _, otherSeat := range seats {
			var potentialSeat int64

			if seat+2 == otherSeat {
				potentialSeat = seat + 1
				foundSeat := false

				for _, notMySeat := range seats {
					if notMySeat == potentialSeat {
						foundSeat = true
						break
					}
				}

				if !foundSeat {
					fmt.Println(potentialSeat)
				}
			}
		}
	}
}

func convertToBinary(row []byte, zero, one byte) string {
	binary := make([]byte, len(row))

	for i, b := range []byte(row) {
		switch b {
		case zero:
			binary[i] = '0'
		case one:
			binary[i] = '1'
		default:
			panic(fmt.Errorf("unexpected character '%b' expected '%b' or '%b'", b, zero, one))
		}
	}

	return string(binary)
}
