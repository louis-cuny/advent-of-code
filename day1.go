package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getPart1(values []string) int {
	sum := 0
	first := true
	previous := 0
	for _, element := range values {
		el, err := strconv.Atoi(element)
		if err != nil {
			continue
		}

		if first {
			first = false
			previous = el
			continue
		}
		if el > previous {
			sum++
		}
		previous = el

	}
	return sum
}

func getPart2(values []string) int {
	sum := 0
	first := true
	second := true
	third := true
	previous := 0
	previous2 := 0
	previous3 := 0
	for _, element := range values {
		el, err := strconv.Atoi(element)
		if err != nil {
			continue
		}

		if first {
			first = false
			previous = el
			continue
		}
		if second {
			second = false
			previous2 = el
			continue
		}
		if third {
			third = false
			previous3 = el
			continue
		}

		if previous+previous2+previous3 < el+previous2+previous3 {
			sum++
		}
		previous = previous2
		previous2 = previous3
		previous3 = el

	}
	return sum
}

func main() {

	filename := "resources/day1.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	values := strings.Split(string(content), "\n")

	part1 := getPart1(values)
	part2 := getPart2(values)

	fmt.Println(part1)
	fmt.Println(part2)
}
