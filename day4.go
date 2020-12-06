package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"regexp"
)

func getMappedFields(passport string) map[string]string {
	fieldMap := make(map[string]string)
	fieldArray := strings.Split(passport, "\n")
	fields := []string {}
	for _, fieldList := range fieldArray {
		fields = append(fields, strings.Split(fieldList, " ")...)
	}
	for _, field := range fields {
		fieldParts := strings.Split(field, ":")
		fieldMap[fieldParts[0]] = fieldParts[1]
	}

	return fieldMap
}


func checkPassPart1(passport string) bool {
	mFields := []string {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	fields := getMappedFields(passport)

	for key, _ := range fields {
		for i, item := range mFields {
			if item == key {
				mFields[i] = mFields[len(mFields)-1]
				mFields[len(mFields)-1] = ""
				mFields = mFields[:len(mFields)-1]
			}
		}

	}

	return len(mFields) == 0
}

func isValid(key string, value string) bool {
	switch key {
	case "byr":
		if len(value) != 4 {
			return false
		}

		year, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		if year >= 1920 && year <= 2002 {
			return true
		}
		return false
	case "iyr":
		if len(value) != 4 {
			return false
		}

		year, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		if year >= 2010 && year <= 2020 {
			return true
		}
		return false
	case "eyr":
		if len(value) != 4 {
			return false
		}

		year, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		if year >= 2020 && year <= 2030 {
			return true
		}
		return false
	case "hgt":
		rValue :=[]rune(value)
		unit := string(rValue[len(rValue)-2:])
		height, err := strconv.Atoi(string(rValue[0 : len(rValue)-2]))
		if err != nil {
			return false
		}

		if unit == "cm" {
			if height >= 150 && height <= 193 {
				return true
			}
		} else if unit == "in" {
			if height >= 59 && height <= 76 {
				return true
			}
		}

		return false
	case "hcl":
		re, err := regexp.MatchString("#[a-f 0-9]{6}", value)
		if err != nil {
			return false
		}

		return re
	case "ecl":
		for _, item := range []string {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if value == item {
				return true
			}
		}

		return false
	case "pid":
		re, err := regexp.MatchString("^[0-9]{9}$", value)
		if err != nil {
			return false
		}

		return re
	default:
		return true
	}
}

func checkPassPart2(passport string) bool {
	if !checkPassPart1(passport) {
		return false
	}

	fields := getMappedFields(passport)
	for key, value := range fields {
		if !isValid(key, value) {
			return false
		}
	}

	return true
}

func main() {
	filename := "resources/day4.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	passports := strings.Split(string(content), "\n\n")

	// Part 1
	nbValidPart1 := 0
	for _, passport := range passports {
		if checkPassPart1(passport) {
			nbValidPart1++
		}
	}

	fmt.Println(nbValidPart1)

	// Part 2
	nbValidPart2 := 0
	for _, passport := range passports {
		if checkPassPart2(passport) {
			nbValidPart2++
		}
	}

	fmt.Println(nbValidPart2)
}
