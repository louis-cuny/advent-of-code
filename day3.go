package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type nextPos func(int, int) int

func countTrees(values []string, getNext nextPos) int {
	pos := 0
	nbTrees := 0
	for index, value := range values {
		if pos > 1000 {
			// Handling special case for line skip.
			pos = getNext(pos - 1100, index)
			continue
		} else if pos < 0 {
			pos = pos + 31
		} else if pos > 30 {
			pos = pos - 31
		}

		if value == "" {
			continue
		}

		if string([]rune(value)[pos]) == "#" {
			nbTrees++
		}

		pos = getNext(pos, index)
	}
	return nbTrees
}

func main() {
	filename := "resources/day3.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	values := strings.Split(string(content), "\n")

	// Part 1
	fmt.Println(countTrees(values, func(i int, vIndex int) int {
		return i + 3
	}))

	// Part 2
	nbTrees := countTrees(values, func(i int, vIndex int) int {
		return i + 1
	}) * countTrees(values, func(i int, vIndex int) int {
		return i + 3
	}) * countTrees(values, func(i int, vIndex int) int {
		return i + 5
	}) * countTrees(values, func(i int, vIndex int) int {
		return i + 7
	}) * countTrees(values, func(i int, vIndex int) int {
		if vIndex % 2 == 0 {
			return i + 1100
		}

		return i + 1
	})

	fmt.Println(nbTrees)
}
