package main

import (
	util "advent-of-code-2023"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type coords = [2]int

var digits = []byte{
	'0',
	'1',
	'2',
	'3',
	'4',
	'5',
	'6',
	'7',
	'8',
	'9',
}

var ignoreChar byte = '.'
var gearChar byte = '*'

func parse(input string) [][]byte {
	lines := strings.Split(input, "\n")
	return util.Map(lines, func(line string) []byte {
		return []byte(line)
	})
}

func isOutOfBounds(input [][]byte, coords coords) bool {
	if coords[0] < 0 {
		return true
	}
	if coords[1] < 0 {
		return true
	}
	if coords[0] >= len(input) {
		return true
	}
	if coords[1] >= len(input[coords[0]]) {
		return true
	}

	return false
}

func valueAt(input [][]byte, coords coords) byte {
	if isOutOfBounds(input, coords) {
		return ignoreChar
	}

	return input[coords[0]][coords[1]]
}

func isDigit(v byte) bool {
	return slices.Contains(digits, v)
}

func hasAdjacentSymbol(input [][]byte, coords coords) bool {
	return slices.ContainsFunc(getAdjacentSymbolsInfo(input, coords), func(info CellInfo) bool {
		return !isDigit(info.CellValue) && info.CellValue != ignoreChar
	})
}

type CellInfo struct {
	CellValue  byte
	CellCoords coords
}

type CellNumberInfo struct {
	IsNumber    bool
	Value       int
	Length      int
	ShouldCount bool
	cellInfo    CellInfo
}

func getAdjacentSymbolsInfo(input [][]byte, coords coords) []CellInfo {

	tlCoords := [2]int{coords[0] + 1, coords[1] - 1} //+1-1 top left
	t_Coords := [2]int{coords[0] + 1, coords[1] + 0} //+1+0 top center
	trCoords := [2]int{coords[0] + 1, coords[1] + 1} //+1+1 top right
	r_Coords := [2]int{coords[0] + 0, coords[1] + 1} //+0+1 middle righ
	brCoords := [2]int{coords[0] - 1, coords[1] + 1} //-1+1 bottom righ
	b_Coords := [2]int{coords[0] - 1, coords[1] - 0} //-1-0 bottom cent
	blCoords := [2]int{coords[0] - 1, coords[1] - 1} //-1-1 bottom left
	l_Coords := [2]int{coords[0] + 0, coords[1] - 1} //+0-1 middle left

	tl := CellInfo{CellValue: valueAt(input, tlCoords), CellCoords: tlCoords} //+1-1 top left
	t_ := CellInfo{CellValue: valueAt(input, t_Coords), CellCoords: t_Coords} //+1+0 top center
	tr := CellInfo{CellValue: valueAt(input, trCoords), CellCoords: trCoords} //+1+1 top right
	r_ := CellInfo{CellValue: valueAt(input, r_Coords), CellCoords: r_Coords} //+0+1 middle right
	br := CellInfo{CellValue: valueAt(input, brCoords), CellCoords: brCoords} //-1+1 bottom right
	b_ := CellInfo{CellValue: valueAt(input, b_Coords), CellCoords: b_Coords} //-1-0 bottom center
	bl := CellInfo{CellValue: valueAt(input, blCoords), CellCoords: blCoords} //-1-1 bottom left
	l_ := CellInfo{CellValue: valueAt(input, l_Coords), CellCoords: l_Coords} //+0-1 middle left

	return []CellInfo{tl, t_, tr, r_, br, b_, bl, l_}
}

func getNumberInfo(input [][]byte, coords coords) CellNumberInfo {
	isNumber := false
	digits := []byte{}
	shouldCount := false
	length := 1
	value := 0

	i := 0

	//if it's the middle of a number, go back to the start.
	for {
		newCoords := [2]int{coords[0], coords[1] + i}
		char := valueAt(input, newCoords)
		if isDigit(char) {
			i--
		} else {
			if i < 0 { //since we alway subtract the index before chenking the next value, we then have to add back one, unless is the first iteration.
				i++
			}
			break
		}
	}

	numStartColIndex := coords[1] + i

	for {
		newCoords := [2]int{coords[0], coords[1] + i}
		char := valueAt(input, newCoords)

		if !isDigit(char) {
			if length > 1 { //since we alway increment the length before chenking the next value, we then have to subtract, unless is the first iteration.
				length--
			}
			break
		}

		if !isNumber {
			isNumber = true
		}

		digits = append(digits, char)

		if !shouldCount && hasAdjacentSymbol(input, newCoords) {
			shouldCount = true
		}
		i++
		length++
	}

	if v, err := strconv.Atoi(string(digits)); err == nil {
		value = v
	}

	cellInfo := CellInfo{CellValue: valueAt(input, coords), CellCoords: [2]int{coords[0], numStartColIndex}}

	return CellNumberInfo{
		IsNumber:    isNumber,
		Value:       value,
		Length:      length,
		ShouldCount: shouldCount,
		cellInfo:    cellInfo,
	}

}

func getGearInfo(input [][]byte, adjacentCells []CellInfo) (bool, int) {
	gearRatio := 1
	isGear := false
	totNums := 0
	numberSet := []coords{} //for tracking duplicates

	for i := 0; i < 8; i++ {
		cell := adjacentCells[i]
		numInfo := getNumberInfo(input, cell.CellCoords)
		if numInfo.IsNumber && !slices.ContainsFunc(numberSet, func(elem [2]int) bool {
			return elem[0] == numInfo.cellInfo.CellCoords[0] && elem[1] == numInfo.cellInfo.CellCoords[1]
		}) {
			gearRatio *= numInfo.Value
			totNums++
			numberSet = append(numberSet, numInfo.cellInfo.CellCoords)
		}
	}

	if totNums == 2 {
		isGear = true
	}

	return isGear, gearRatio
}

func part_1(input [][]byte) int {
	sum := 0
	for rowIndex, line := range input {
		for colIndex := 0; colIndex < len(line); colIndex++ {
			coords := [2]int{rowIndex, colIndex}
			numberInfo := getNumberInfo(input, coords)

			if !numberInfo.IsNumber {
				continue
			}

			colIndex += numberInfo.Length - 1 //skip the rest of the number

			if numberInfo.ShouldCount {
				sum += numberInfo.Value
			}
		}
	}
	return sum
}

func part_2(input [][]byte) int {
	sum := 0
	for rowIndex, line := range input {
		for colIndex := 0; colIndex < len(line); colIndex++ {
			coords := [2]int{rowIndex, colIndex}
			char := valueAt(input, coords)
			if char != gearChar {
				continue
			}

			adjacentInfo := getAdjacentSymbolsInfo(input, coords)
			isGear, gearRatio := getGearInfo(input, adjacentInfo)
			if isGear {
				sum += gearRatio
			}
		}
	}
	return sum
}

func main() {
	content, _ := os.ReadFile("day3/input")
	parsedContent := parse(string(content))
	fmt.Println("part1: ", part_1(parsedContent))
	fmt.Println("part2: ", part_2(parsedContent))
}
