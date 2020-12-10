package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var bags []string
var bags2 []string

func contains(items []string, string string) bool {
	for _, item := range items {
		if item == string {
			return true
		}
	}

	return false
}

func stringInArray(list []string, string string) bool {
	for _, value := range list {
		if value == string {
			return true
		}
	}

	return false
}

func part1(lines []string, search string) {
	for _, line := range lines {
		items := regexp.MustCompile(" bags contain [0-9] | bag, [0-9] | bags, [0-9] | bags\\.| bag\\.| bags contain no other bags\\.| bag contain no other bags\\.").Split(line, -1)

		if contains(items, search) && !stringInArray(bags, items[0]) {
			bags = append(bags, items[0])
			part1(lines, items[0])
		}
	}
}

func part2(lines []string, search string) int {
	fa := 0
	for _, line := range lines {
		items := strings.Split(line, " bags contain ")
		if items[0] != search {
			continue
		}

		items = regexp.MustCompile(" bag, | bags, | bags\\.| bag\\.").Split(items[1], -1)
		items = items[:len(items)-1]
		for _, item := range items {
			if len(item) == 0 {
				continue
			}

			f, _ := strconv.Atoi(string(item[0]))
			fa += f + f*part2(lines, item[2:])
		}
	}

	return fa
}

func main() {
	filename := "resources/day7.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	lines := strings.Split(string(content), "\n")

	// Skipping useless lines.
	var fLines []string
	for _, line := range lines {
		if !strings.Contains(line, "contain no other bags.") {
			fLines = append(fLines, line)
		}
	}

	part1(lines, "shiny gold")
	fmt.Println(len(bags) - 1)
	fmt.Println(part2(fLines, "shiny gold"))
}
