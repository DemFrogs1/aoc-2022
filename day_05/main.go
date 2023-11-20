package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/DemFrogs1/aoc-2022/lib"
)

func main() {
	data := lib.Parse("input.txt")

	stack, rearrangements := ParseStacks(data)

	fmt.Println(stack)

	for _, rearg := range rearrangements {
		rearrangement := ParseRearrangement(rearg)
		stack = Rearrange(stack, rearrangement)
	}
	fmt.Println(stack)
}

func Rearrange(stack map[string][]string, rearrangement []int) map[string][]string {
	count := rearrangement[0]
	fromKey := strconv.Itoa(rearrangement[1])
	toKey := strconv.Itoa(rearrangement[2])

	from := stack[fromKey]
	to := stack[toKey]

	if len(from) >= count {
		newFrom := make([]string, count)
		copy(newFrom, from[:count])

		//reverse the array for part 1, don't reverse for part 2
		for i, j := 0, len(newFrom)-1; i < j; i, j = i+1, j-1 {
			newFrom[i], newFrom[j] = newFrom[j], newFrom[i]
		}

		stack[toKey] = append(newFrom, to...)

		stack[fromKey] = from[count:]
		if fromKey == "8" && toKey == "7" && count == 9 {
			fmt.Println(stack)
		}
	}

	return stack
}
func ParseStacks(data []string) (map[string][]string, []string) {
	stacks, stacksNumber, rearrangements := GetStackInfo(data)

	finalStack := make(map[string][]string)

	for _, num := range stacksNumber {
		number := string(num)
		if strings.TrimSpace(number) != "" {
			numberIndex := strings.Index(stacksNumber, number)

			for _, stack := range stacks[:len(stacks)-1] {
				crate := stack[numberIndex]

				if strings.TrimSpace(string(crate)) != "" {
					finalStack[number] = append(finalStack[number], string(crate))
				}
			}
		}
	}

	return finalStack, rearrangements
}
func ParseRearrangement(rearrangement string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(rearrangement, -1)

	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}
func GetStackInfo(data []string) ([]string, string, []string) {
	var stacks []string
	var rearrangements []string

	iteratedOverStack := false
	for _, row := range data {
		if strings.TrimSpace(row) == "" {
			iteratedOverStack = true
		} else {
			if !iteratedOverStack {
				stacks = append(stacks, row)
			} else {
				rearrangements = append(rearrangements, row)
			}
		}
	}

	stacksNumber := stacks[len(stacks)-1]

	return stacks, stacksNumber, rearrangements
}
