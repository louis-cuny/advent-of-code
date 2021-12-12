package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func getPart1D7(crabs []int) int {
	sort.Ints(crabs)
	med := crabs[len(crabs)/2]
	fuel := 0
	for _, crab := range crabs {
		if med > crab {
			fuel += med - crab
		} else {
			fuel += crab - med
		}
	}

	return fuel
}

func getPart2D7(crabs []int) int {
	sum := 0
	for _, crab := range crabs {
		sum += crab
	}

	avg := sum / len(crabs)
	fuel := 0
	for _, crab := range crabs {
		fuel += calcFuel(avg, crab)
	}

	return fuel
}

func calcFuel(target int, start int) int {
	nb := 0
	if target > start {
		nb = target - start
	} else {
		nb = start - target
	}
	sum := 0
	for i := 1; i <= nb; i++ {
		sum += i
	}

	return sum
}

func main() {
	filename := "resources/day7.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	var crabs []int
	rawCrabs := strings.Split(string(content), ",")
	for _, rawCrab := range rawCrabs {
		value, _ := strconv.Atoi(rawCrab)
		crabs = append(crabs, value)
	}
	part1 := getPart1D7(crabs)
	part2 := getPart2D7(crabs)
	fmt.Println(part1)
	fmt.Println(part2)
}
