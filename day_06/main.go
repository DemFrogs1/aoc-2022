package main

import (
	"fmt"
	"strings"

	"github.com/DemFrogs1/aoc-2022/lib"
)

func main() {
	data := lib.Parse("input.txt")

	//Change 4 to 14 to get part 2 answer
	fmt.Println(FindPacket(data, 0, 4))
}

func FindPacket(data []string, startString, endString int) int {
	str := ""
	maxChar := endString

outerloop:
	for {
		for index, char := range data[0][startString:endString] {

			if !strings.ContainsRune(str, char) {
				str += string(char)
			}

			if index == maxChar-1 {
				if len(str) == maxChar {
					break outerloop
				}

				str = ""
			}
		}
		startString++
		endString++
	}

	return endString
}
