package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/DemFrogs1/aoc-2022/lib"
)

const sizeThreshold = 100000
const fileSystemSize = 70000000
const requiredSpaceForUpdate = 30000000

type FileSystem struct {
	Structure map[string][]string
}

func main() {
	data := lib.Parse("input.txt")
	fileSystem := NewFileSystem()

	currKey := ""
	subCategories := make(map[string]string)

	for _, line := range data {
		if strings.HasPrefix(line, "$ cd") {
			currKey = handleCdCommand(line, &fileSystem, subCategories)
			continue
		}

		if line[0:1] == "$" || currKey == "/" {
			continue
		}

		handleFileOrDir(line, currKey, &fileSystem, subCategories)
	}

	totalSum := 0
	usedSpace := 0
	for key, value := range fileSystem.Structure {
		sum := calculateSum(value, &fileSystem.Structure)
		if sum <= sizeThreshold {
			totalSum += sum
		}
		if subCategories[key] == "" {
			usedSpace += sum
		}
	}

	fmt.Printf("The part 1 answer is %d\n", totalSum)
	sizeToDelete := part2(&fileSystem, usedSpace)
	fmt.Printf("The part 2 answer is %d\n", sizeToDelete)
}

func part2(fileSystem *FileSystem, usedSpace int) int {
	smallestDir := math.MaxInt
	for _, value := range fileSystem.Structure {
		sum := calculateSum(value, &fileSystem.Structure)
		if sum > requiredSpaceForUpdate-(fileSystemSize-usedSpace) && sum < smallestDir {
			smallestDir = sum
		}
	}
	return smallestDir
}

func handleCdCommand(line string, fileSystem *FileSystem, subCategories map[string]string) string {
	currKey := strings.TrimSpace(line[4:])
	if currKey == ".." || currKey == "/" {
		return currKey
	}

	id := checkRedundancy(currKey, fileSystem.Structure)
	if id > 0 {
		currKey = currKey + "_" + strconv.Itoa(id+1)
	}
	fileSystem.Structure[currKey] = make([]string, 0)

	return currKey
}

func handleFileOrDir(line, currKey string, fileSystem *FileSystem, subCategories map[string]string) {
	if !strings.Contains(line, "dir") {
		line = strings.Split(line, " ")[0]
	} else {
		dir := strings.Split(line, " ")[1]
		id := checkRedundancy(dir, fileSystem.Structure)
		if id > 0 {
			dir = dir + "_" + strconv.Itoa(id+1)
		}
		subCategories[dir] = currKey
		line = "dir " + dir
	}

	fileSystem.Structure[currKey] = append(fileSystem.Structure[currKey], line)
}

func calculateSum(sizes []string, fileSystem *map[string][]string) int {
	sum := 0
	for _, file := range sizes {
		if strings.HasPrefix(file, "dir") {
			dir := strings.Split(file, " ")[1]
			sum += calculateSum((*fileSystem)[dir], fileSystem)
		} else {
			num, err := strconv.Atoi(file)
			if err != nil {
				log.Println(err)
				continue
			}
			sum += num
		}
	}

	return sum
}

func checkRedundancy(str string, fileSystem map[string][]string) int {
	count := 0
	for key := range fileSystem {
		if strings.HasPrefix(key, str) {
			count++
		}
	}
	return count
}

func NewFileSystem() FileSystem {
	return FileSystem{
		Structure: make(map[string][]string),
	}
}
