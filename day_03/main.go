package main

import (
	"fmt"
	"strings"

	"github.com/DemFrogs1/aoc-2022/lib"
)

var allLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var groups []string

func main() {
	data := lib.Parse("input.txt")

	secondPartTotal := 0
	firstPartTotal := 0

	for _, entry := range data {
		firstPartTotal += part1(entry)

		secondPartTotal += part2(entry)
	}

	fmt.Printf("The total priority is %d and part 1 is %d", secondPartTotal, firstPartTotal)
}

func part1(entry string) int {
	entryMiddle := len(entry) / 2

	firstEntry := entry[0:entryMiddle]
	secondEntry := entry[entryMiddle:]

	commonLetter := GetCommonFromStrings(firstEntry, secondEntry)

	return strings.IndexRune(allLetters, commonLetter) + 1
}

func part2(entry string) int {
	groups = append(groups, entry)

	if len(groups) == 3 {
		commonLetter := GetCommonFromStrings(groups[0], groups[1], groups[2])
		groups = nil
		return strings.IndexRune(allLetters, commonLetter) + 1
	} else {
		return 0
	}
}

func GetCommonFromStrings(entries ...string) rune {
	if len(entries) < 2 {
		return 0
	}

	var commonLetter rune

	for _, char := range entries[0] {
		foundInAll := true

		for _, str := range entries[1:] {
			if !strings.ContainsRune(str, char) {
				foundInAll = false
				break
			}
		}

		if foundInAll {
			commonLetter = char
			break
		}
	}

	return commonLetter
}
