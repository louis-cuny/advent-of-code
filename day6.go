package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func countAnswersPart1(group string) int {
	res := make(map[string]int)
	for _, answers := range strings.Split(group, "\n") {
		for _, answer := range answers {
			res[string(answer)] = res[string(answer)] + 1
		}
	}

	return len(res)
}

func countAnswersPart2(group string) int {
	res := make(map[string]int)
	lines := strings.Split(group, "\n")
	sum := 0
	nbLines := len(lines)
	for _, answers := range lines {
		for _, answer := range answers {
			res[string(answer)] = res[string(answer)] + 1
			if nbLines == res[string(answer)] {
				sum++
			}
		}
	}

	return sum
}

func main() {
	filename := "resources/day6.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	groups := strings.Split(string(content), "\n\n")

	// Part 1
	sum := 0
	for _, group := range groups {
		sum += countAnswersPart1(group)
	}

	fmt.Println(sum)

	// Part 2
	sum = 0
	for _, group := range groups {
		sum += countAnswersPart2(group)
	}

	fmt.Println(sum)
}
