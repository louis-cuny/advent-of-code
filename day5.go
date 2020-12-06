package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func getRow(codes string) int {
	lFront := 0
	lBack := 127

	for _, code := range []rune(codes) {
		switch string(code) {
		case "F":
			lBack = lFront + (lBack-lFront)/2
			break
		case "B":
			lFront = int(math.Ceil(float64(lFront) + (float64(lBack-lFront))/2))
			break
		}
	}

	if lFront != lBack {
		fmt.Println("We got a problem.")
	}

	return lFront
}

func getCol(codes string) int {
	lLeft := 0
	lRight := 7

	for _, code := range []rune(codes) {
		switch string(code) {
		case "L":
			lRight = lLeft + (lRight-lLeft)/2
			break
		case "R":
			lLeft = int(math.Ceil(float64(lLeft) + (float64(lRight-lLeft))/2))
			break
		}
	}

	if lRight != lLeft {
		fmt.Println("We got a problem.")
	}

	return lLeft
}

func main() {
	filename := "resources/day5.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	seatCodes := strings.Split(string(content), "\n")

	// Part 1
	highest := 0
	highestV2 := 0
	list := []int{}
	for _, seatCode := range seatCodes {
		seatRowCode := seatCode[:len(seatCode)-3]

		seatColCode := seatCode[len(seatCode)-3:]

		seatRow := getRow(seatRowCode)
		seatCol := getCol(seatColCode)

		if seatRow*8+seatCol > highest {
			highest = seatRow*8 + seatCol
		}

		seatRowCode = strings.ReplaceAll(seatRowCode, "F", "0")
		seatRowCode = strings.ReplaceAll(seatRowCode, "B", "1")
		seatColCode = strings.ReplaceAll(seatColCode, "L", "0")
		seatColCode = strings.ReplaceAll(seatColCode, "R", "1")

		seatRowV2 := getNum(seatRowCode)
		seatColV2 := getNum(seatColCode)

		if seatRowV2*8+seatColV2 > highestV2 {
			highestV2 = seatRowV2*8 + seatColV2
		}

		list = append(list, seatRowV2*8+seatColV2)
	}

	if highest == highestV2 {
		fmt.Println(highest)
	}

	// Part 2
	mySeat := 0
	for seatNum := 0; seatNum < highestV2; seatNum++ {
		if inArray(list, seatNum) {
			continue
		}

		foundPrevious := false
		for _, value := range list {
			if value == seatNum-1 {
				foundPrevious = true
				break
			}
		}
		foundNext := false
		for _, value := range list {
			if value == seatNum+1 {
				foundNext = true
				break
			}
		}

		if foundPrevious && foundNext {
			mySeat = seatNum
			break
		}
	}

	fmt.Println(mySeat)
}

func getNum(code string) int {
	output, err := strconv.ParseInt(code, 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	return int(output)
}

func inArray(list []int, seatNum int) bool {
	for _, value := range list {
		if value == seatNum  {
			return true
		}
	}

	return false
}
