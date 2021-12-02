package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


func main() {
	filename := "resources/day12.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	lines := strings.Split(string(content), "\n")

	var instructions map[rune]int
	for i, line := range lines {
		instructions
	}

	//fmt.Println(d12Part1(array))
	//fmt.Println(d12Part2(array))
}
