package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

var threesIds []int

func d10Part1(numbers []int) int {
	ones := 1
	threes := 1
	for i, number := range numbers {
		if i == 0 {
			continue
		}

		if number-numbers[i-1] == 1 {
			ones++
		} else if number-numbers[i-1] == 3 {
			threes++
		} else {
			return -1
		}
	}

	return ones * threes
}

func collectThreesIds(numbers []int) {
	for i, number := range numbers {
		if i == 0 {
			continue
		}
		if number-numbers[i-1] == 3 {
			threesIds = append(threesIds, i)
		}
	}
}


func d10Part2(numbers []int) int64 {
	var count int64 = 1

	i := 0
	for _, id := range threesIds {
		if id-i < 3 {
			i = id
			continue
		}

		pw := int64(math.Pow(2, float64(id-i-2)))
		diff := id-i-4
		if diff > 0 {
			pw -= int64(diff)
		}

		count *= pw
		i = id
	}

	return count
}

func main() {
	filename := "resources/day10.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	lines := strings.Split(string(content), "\n")
	var numbers []int
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}

	sort.Ints(numbers)
	d10part1 := d10Part1(numbers)
	fmt.Println(d10part1)

	// Adding the 0 and target.
	numbers = append([]int{0}, numbers...)
	numbers = append(numbers, numbers[len(numbers)-1] + 3)

	collectThreesIds(numbers)
	d10part2 := d10Part2(numbers)
	fmt.Println(d10part2)

}
