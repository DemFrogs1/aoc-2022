package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DemFrogs1/aoc-2022/lib"
)

func main() {
	data := lib.Parse("input.txt")
	collusionCount := 0
	overlapCount := 0
	for _, pair := range data {
		res := strings.Split(pair, ",")
		var formattedPair [][]int
		for _, entry := range res {
			result, err := GetFirstAndLast(entry)
			if err != nil {
				log.Fatalf("conversion error")
			}

			formattedPair = append(formattedPair, result)
		}
		if len(formattedPair) > 2 {
			formattedPair = nil
		}
		if CheckCollusion(formattedPair) {
			collusionCount++
		}
		if CheckOverlap(formattedPair) {
			overlapCount++
		}
	}
	fmt.Println(collusionCount, overlapCount)
}

func GetFirstAndLast(entry string) ([]int, error) {
	res := strings.Split(entry, "-")
	var resultInt []int

	for _, str := range res {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		resultInt = append(resultInt, num)
	}

	return resultInt, nil
}

// part 1
func CheckCollusion(pair [][]int) bool {
	firstPair := pair[0]
	secondPair := pair[1]

	if (firstPair[0] <= secondPair[0] && firstPair[1] >= secondPair[1]) ||
		(firstPair[0] >= secondPair[0] && firstPair[1] <= secondPair[1]) {
		return true
	}
	return false
}

// part 2
func CheckOverlap(pair [][]int) bool {
	firstPair := pair[0]
	secondPair := pair[1]

	if (firstPair[0] < secondPair[0] && firstPair[1] < secondPair[0]) ||
		(secondPair[0] < firstPair[0] && secondPair[1] < firstPair[0]) {
		return false
	}

	return true
}
