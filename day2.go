package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getPart1D2(values []string) int {
	depth := 0
	horizontal := 0
	for _, element := range values {
		num, err := strconv.Atoi(element[len(element)-1:])
		if err != nil {
			continue
		}

		switch element[:1] {
		case "f":
			horizontal += num
		case "u":
			depth -= num
		case "d":
			depth += num
		}
	}

	return depth * horizontal
}

func getPart2D2(values []string) int {
	depth := 0
	horizontal := 0
	aim := 0
	for _, element := range values {
		num, err := strconv.Atoi(element[len(element)-1:])
		if err != nil {
			continue
		}

		switch element[:1] {
		case "f":
			horizontal += num
			depth += num * aim
		case "u":
			aim -= num
		case "d":
			aim += num
		}
	}

	return depth * horizontal
}

func main() {

	filename := "resources/day2.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	values := strings.Split(string(content), "\n")

	part1 := getPart1D2(values)
	part2 := getPart2D2(values)

	fmt.Println(part1)
	fmt.Println(part2)
}
