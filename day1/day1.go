package main

import (
	util "advent-of-code-2023"
	"os"
	"strings"
)

/*

example input:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

goal:

for each line concatenate the first digit with the last digit to get a two digit **number**,
then sum all the numbers together

*/

// takes in a string that must have at least on digit in it.
// returns the first and last digit concatenated together as a number]

var digitMap []string = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func valueOfDigit(digit string) int {
	for index, value := range digitMap {
		if value == digit && index < 10 {
			return index
		} else if value == digit {
			return index - 10
		}
	}
	return -1
}

func parseLine(line string) []string {
	result := []string{}
	for index, _ := range line {
		for _, digit := range digitMap {
			if index+len(digit) > len(line) {
				continue
			}
			if line[index:index+len(digit)] == digit {
				result = append(result, digit)
				break
			}
		}
	}

	return result
}

func parse(input string) [][]string {
	var lines []string = strings.Split(input, "\n")
	result := [][]string{}
	for _, line := range lines {
		result = append(result, parseLine(line))
	}
	return result
}

func part_1(input string) int {
	var sum int = 0
	var data = parse(input)
	for _, line := range data {
		var firstDigit = util.First(line, func(digit string) bool {
			return len(digit) == 1
		})
		var lastDigit = util.Last(line, func(digit string) bool {
			return len(digit) == 1
		})

		sum += valueOfDigit(firstDigit)*10 + valueOfDigit(lastDigit)
	}

	return sum
}

func part_2(input string) int {
	var sum int = 0
	var data = parse(input)

	for _, line := range data {
		var firstDigit = line[0]
		var lastDigit = line[len(line)-1]
		sum += valueOfDigit(firstDigit)*10 + valueOfDigit(lastDigit)
	}

	return sum
}

func main() {
	if content, err := os.ReadFile("day1/input"); err == nil {
		var trimmedString = strings.TrimSpace(string(content))
		println("part_1", part_1(trimmedString))
		println("part_2", part_2(trimmedString))
	} else {
		panic(err)
	}
}
