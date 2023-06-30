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
	firstResult := 0
	trueResult := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		enemy := int(line[0] - 'A')
		player := int(line[2] - 'X')

		firstStratMatchResult := 1 + player + ((player-enemy+4)%3)*3
		firstResult += firstStratMatchResult

		trueStratMatchResult := player*3 + (enemy+(player-1)+3)%3 + 1
		trueResult += trueStratMatchResult
	}

	fmt.Printf("First result: %d\n", firstResult)
	fmt.Printf("True result: %d\n", trueResult)
}
