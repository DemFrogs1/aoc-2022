package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/DemFrogs1/aoc-2022/lib"
)

func main() {
	data := lib.Parse("input.txt")

	rows := len(data[0])
	col := len(data)
	matrix := make([][]int, col)
	for i := range matrix {
		matrix[i] = make([]int, rows)
	}

	for dataIndex, line := range data {
		for lineIndex, num := range line {
			n, err := strconv.Atoi(string(num))
			if err == nil {
				matrix[dataIndex][lineIndex] = n
			}
		}
	}

	visible, scenicScores := countVisibleTrees(matrix)
	edges := (col * 2) + ((col - 2) * 2)

	fmt.Println(visible+edges, slices.Max(scenicScores))
}

func countVisibleTrees(matrix [][]int) (int, []int) {
	visbileCount := 0

	var scenicScores []int

	for index, row := range matrix[1 : len(matrix)-1] {
		rowTrueIndex := index + 1
		for index, data := range row[1 : len(row)-1] {
			trueIndex := index + 1
			isVisible, scenicScore := compare(matrix, rowTrueIndex, data, trueIndex)
			scenicScores = append(scenicScores, scenicScore)
			if isVisible {
				visbileCount++
			}
		}
	}

	return visbileCount, scenicScores
}

func compare(matrix [][]int, valueRowIndex, treeSize, valueIndex int) (bool, int) {
	isVisible := false

	treeProx := getTreeProximity(matrix, valueRowIndex, valueIndex)

	scenicScore := 1
	for key, value := range treeProx {
		allTreesLarger := false
		sideScenicScore := 0

		start, end, step := 0, len(value), 1
		reverse := key == "left" || key == "top"
		if reverse {
			start, end, step = len(value)-1, -1, -1
		}

		for i := start; i != end; i += step {
			prox := value[i]
			sideScenicScore++
			if prox >= treeSize {
				allTreesLarger = true
				break
			}
		}
		if !allTreesLarger {
			isVisible = true
		}

		scenicScore *= sideScenicScore
	}

	return isVisible, scenicScore
}

func getTreeProximity(matrix [][]int, valueRowIndex, valueIndex int) map[string][]int {
	treeProx := map[string][]int{
		"top":    []int{},
		"bottom": []int{},
		"right":  []int{},
		"left":   []int{},
	}
	treeProx["right"] = append(treeProx["right"], matrix[valueRowIndex][1+valueIndex:]...)
	treeProx["left"] = append(treeProx["left"], matrix[valueRowIndex][:valueIndex]...)

	for _, row := range matrix[:valueRowIndex] {
		if valueIndex < len(row) {
			treeProx["top"] = append(treeProx["top"], row[valueIndex])
		}
	}
	for _, row := range matrix[1+valueRowIndex:] {
		if valueIndex < len(row) {
			treeProx["bottom"] = append(treeProx["bottom"], row[valueIndex])
		}
	}

	return treeProx
}
