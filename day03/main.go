package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.data")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	// El plan es hacer un array de tamaño 26*2 bool falsos, rellenar con true la primera mitad y luego chequear la segunda para encontrar el error.
	repeatedSum := 0
	common3Sum := 0
	subgroupIdx := 0
	var countGroupLetters map[byte]int
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if subgroupIdx == 0 {
			countGroupLetters = make(map[byte]int)
		}
		foundLetters := make(map[byte]int)
		// First part of the problem
		for i := 0; i < (len(line) / 2); i++ {
			foundLetters[line[i]]++
		}
		for i := (len(line) / 2); i < len(line); i++ {
			if foundLetters[line[i]] != 0 {
				repeatedSum += letterToValue(line[i])
				break
			}
		}
		// Second part of the problem
		for i := 0; i < len(line); i++ {
			if countGroupLetters[line[i]] == subgroupIdx {
				countGroupLetters[line[i]]++
			}
			if countGroupLetters[line[i]] == 3 {
				common3Sum += letterToValue(line[i])
				break
			}
		}
		subgroupIdx++
		if subgroupIdx > 2 {
			subgroupIdx = 0
		}
	}
	fmt.Printf("Total repeatedSum: %d\n", repeatedSum)
	fmt.Printf("Total common3Sum: %d\n", common3Sum)
}

func letterToValue(letter byte) int {
	if 'a' <= letter && letter <= 'z' {
		return int(letter-'a') + 1
	} else {
		return int(letter-'A') + 27
	}
}
