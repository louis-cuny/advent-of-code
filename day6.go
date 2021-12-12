package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getPart1D6(fishes []int) int {
	for i := 0; i < 80; i++ {
		var newFishes []int
		for _, fish := range fishes {
			if fish == 0 {
				newFishes = append(newFishes, 8, 6)
			} else {
				newFishes = append(newFishes, fish-1)
			}
		}
		fishes = newFishes
	}

	return len(fishes)
}

func getPart2D6(fishes []int) int {
	var groups = make(map[int]int)
	for _, fish := range fishes {
		if _, ok := groups[fish]; ok {
			groups[fish]++
		} else {
			groups[fish] = 1
		}
	}

	for i := 0; i < 256; i++ {
		var newGroups = make(map[int]int)
		for val, fac := range groups {
			if val == 0 {
				newGroups[8] = fac
				if _, ok := newGroups[6]; ok {
					newGroups[6] += fac
				} else {
					newGroups[6] = fac
				}
			} else {
				if _, ok := newGroups[6]; ok {
					newGroups[val-1] += fac
				} else {
					newGroups[val-1] = fac
				}
			}
		}
		groups = newGroups
	}

	sum := 0
	for _, group := range groups {
		sum += group
	}

	return sum
}

func main() {
	filename := "resources/day6.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	var fishes []int
	rawFishes := strings.Split(string(content), ",")
	for _, rawFish := range rawFishes {
		value, _ := strconv.Atoi(rawFish)
		fishes = append(fishes, value)
	}
	part1 := getPart1D6(fishes)
	part2 := getPart2D6(fishes)
	fmt.Println(part1)
	fmt.Println(part2)
}
