package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func countAdjacent(value rune, array customArray, y int, x int) int {
	count := 0

	yMax := len(array) - 1
	xMax := len(array[0]) - 1

	// tl
	if x > 0 && y > 0 {
		if array[y-1][x-1] == value {
			count++
		}
	}

	// t
	if y > 0 {
		if array[y-1][x] == value {
			count++
		}
	}

	// tr
	if x < xMax && y > 0 {
		if array[y-1][x+1] == value {
			count++
		}
	}

	// r
	if x < xMax {
		if array[y][x+1] == value {
			count++
		}
	}

	// br
	if x < xMax && y < yMax {
		if array[y+1][x+1] == value {
			count++
		}
	}

	// b
	if y < yMax {
		if array[y+1][x] == value {
			count++
		}
	}

	// bl
	if x > 0 && y < yMax {
		if array[y+1][x-1] == value {
			count++
		}
	}

	// l
	if x > 0 {
		if array[y][x-1] == value {
			count++
		}
	}

	return count
}

func countAdjacent2(value rune, array customArray, y int, x int) int {
	count := 0

	yMax := len(array) - 1
	xMax := len(array[0]) - 1

	// tl
	tly := y
	tlx := x
	for tlx > 0 && tly > 0 {
		tly--
		tlx--
		if array[tly][tlx] == value {
			count++
			break
		}

		if array[tly][tlx] != '.' {
			break
		}
	}

	// t
	ty := y
	for ty > 0 {
		ty--
		if array[ty][x] == value {
			count++
			break
		}

		if array[ty][x] != '.' {
			break
		}
	}

	// tr
	try := y
	trx := x
	for trx < xMax && try > 0 {
		try--
		trx++
		if array[try][trx] == value {
			count++
			break
		}

		if array[try][trx] != '.' {
			break
		}
	}

	// r
	rx := x
	for rx < xMax {
		rx++
		if array[y][rx] == value {
			count++
			break
		}

		if array[y][rx] != '.' {
			break
		}
	}

	// br
	bry := y
	brx := x
	for brx < xMax && bry < yMax {
		bry++
		brx++
		if array[bry][brx] == value {
			count++
			break
		}

		if array[bry][brx] != '.' {
			break
		}
	}

	// b
	by := y
	for by < yMax {
		by++
		if array[by][x] == value {
			count++
			break
		}

		if array[by][x] != '.' {
			break
		}
	}

	// bl
	bly := y
	blx := x
	for blx > 0 && bly < yMax {
		bly++
		blx--
		if array[bly][blx] == value {
			count++
			break
		}

		if array[bly][blx] != '.' {
			break
		}
	}

	// l
	lx := x
	for lx > 0 {
		lx--
		if array[y][lx] == value {
			count++
			break
		}

		if array[y][lx] != '.' {
			break
		}
	}

	return count
}

type customArray [97][90]rune
type counter func(rune, customArray, int, int) int

func getNextArray(array customArray, handler counter, limit int) customArray {
	var nArray customArray
	for y, runes := range array {
		for x, value := range runes {
			if value == '.' {
				nArray[y][x] = '.'
				continue
			}

			if value == 'L' {
				if handler('#', array, y, x) == 0 {
					nArray[y][x] = '#'
				} else {
					nArray[y][x] = 'L'
				}

				continue
			}

			if value == '#' {
				if handler('#', array, y, x) >= limit {
					nArray[y][x] = 'L'
				} else {
					nArray[y][x] = '#'
				}

				continue
			}

		}
	}
	return nArray
}

func countOccupied(array customArray) int {
	count := 0

	for _, runes := range array {
		for _, value := range runes {
			if value == '#' {
				count++
			}
		}
	}

	return count
}

func d11Part1(array customArray) int {
	pArray := array
	for {
		array = getNextArray(array, countAdjacent, 4)
		if pArray == array {
			return countOccupied(array)

		}
		pArray = array
	}
}

func d11Part2(array customArray) int {
	pArray := array
	for {
		array = getNextArray(array, countAdjacent2, 5)
		if pArray == array {
			return countOccupied(array)

		}
		pArray = array
	}
}

func main() {
	filename := "resources/day11.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	lines := strings.Split(string(content), "\n")

	var array customArray
	for i, line := range lines {
		for f, value := range line {
			array[i][f] = value
		}
	}

	fmt.Println(d11Part1(array))
	fmt.Println(d11Part2(array))
}
