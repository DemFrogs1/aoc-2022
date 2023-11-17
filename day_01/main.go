package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/DemFrogs1/aoc-2022/lib"
)

const TOP_MAX = 3

func main() {
	var calories []int
	sum := 0

	data := lib.Parse("input.txt")

	for _, cal := range data {
		if len(cal) > 0 {
			n, err := strconv.Atoi(cal)
			if err != nil {
				log.Println(err)
				continue
			}
			sum += n
			calories = append(calories, sum)
		} else {
			sum = 0
		}
	}

	max := part1(calories)
	totalMax := part2(calories)

	fmt.Printf("max is equal to %d total top 3 is equal to %d", max, totalMax)
}

func calculateMax(numbers []int) (int, int) {
	max := 0
	index := 0

	for i, n := range numbers {
		if max < n {
			max = n
			index = i
		}
	}
	return max, index
}

func part1(calories []int) int {
	max, _ := calculateMax(calories)
	return max
}
func part2(calories []int) int {
	totalMax := 0

	for i := 0; i < TOP_MAX; i++ {
		max, i := calculateMax(calories)
		totalMax += max
		calories = lib.DeleteElement(calories, i)
	}
	return totalMax
}
