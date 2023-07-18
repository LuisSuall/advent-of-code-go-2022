package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func main() {
	file, err := os.Open("input.data")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var nestedZones int64
	var overlapZones int64
	var wg sync.WaitGroup

	for fileScanner.Scan() {
		line := fileScanner.Text()
		wg.Add(1)
		go func() {
			defer wg.Done()
			workZones := strings.Split(line, ",")
			firstElfZones := strings.Split(workZones[0], "-")
			secondElfZones := strings.Split(workZones[1], "-")

			firstElfZonesInt := [2]int{strToInt(firstElfZones[0]), strToInt(firstElfZones[1])}
			secondElfZonesInt := [2]int{strToInt(secondElfZones[0]), strToInt(secondElfZones[1])}

			nested := (secondElfZonesInt[0] <= firstElfZonesInt[0] && firstElfZonesInt[1] <= secondElfZonesInt[1]) ||
				(firstElfZonesInt[0] <= secondElfZonesInt[0] && secondElfZonesInt[1] <= firstElfZonesInt[1])
			if nested {
				atomic.AddInt64(&nestedZones, 1)
				// fmt.Printf("%d-%d in %d-%d\n", firstElfZonesInt[0], firstElfZonesInt[1], secondElfZonesInt[0], secondElfZonesInt[1])
			}

			overlap := (firstElfZonesInt[0] <= secondElfZonesInt[0] && secondElfZonesInt[0] <= firstElfZonesInt[1]) ||
				(secondElfZonesInt[0] <= firstElfZonesInt[0] && firstElfZonesInt[0] <= secondElfZonesInt[1])
			if overlap {
				atomic.AddInt64(&overlapZones, 1)
			}
		}()
	}
	wg.Wait()

	fmt.Printf("Total nested zones: %d\n", nestedZones)
	fmt.Printf("Total overlap zones: %d\n", overlapZones)
}

func strToInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("PANIC!!")
		return 0
	}
	return v
}
