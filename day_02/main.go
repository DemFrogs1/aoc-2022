package main

import (
	"fmt"
	"strings"

	"github.com/DemFrogs1/aoc-2022/lib"
)

type GameRules struct {
	Equivalent string
	point      int
	winsOver   string
	loseTo     string //part 2
	indicates  string //part 2
}

var GameValues = map[string]GameRules{
	"A": {"R", 1, "Z", "Y", ""},
	"X": {"R", 1, "C", "Y", "lose"},

	"B": {"P", 2, "X", "Z", ""},
	"Y": {"P", 2, "A", "Z", "draw"},

	"C": {"S", 3, "Y", "X", ""},
	"Z": {"S", 3, "B", "X", "win"},
}

var GameScores = map[string]int{
	"win":  6,
	"draw": 3,
}

func main() {
	data := lib.Parse("input.txt")
	score := 0
	for _, round := range data {
		res := strings.Split(round, " ")

		elfPlay := GameValues[res[0]]
		userPlay := GameValues[res[1]]

		// part 2 start
		if userPlay.indicates == "draw" {
			userPlay = elfPlay
		} else if userPlay.indicates == "win" {
			userPlay = GameValues[elfPlay.loseTo]
		} else {
			userPlay = GameValues[elfPlay.winsOver]
		}
		//part 2 end

		score += userPlay.point

		if elfPlay.Equivalent == userPlay.Equivalent {
			score += GameScores["draw"]
			continue
		}

		if GameValues[userPlay.winsOver].Equivalent == elfPlay.Equivalent {
			score += GameScores["win"]
			continue
		}
	}
	fmt.Printf("The total score is %d", score)
}
