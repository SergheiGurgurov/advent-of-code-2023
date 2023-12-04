package main

import (
	util "advent-of-code-2023"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadFileSync("day4/input")
	data := parse(&input)
	fmt.Println("part_1", part_1(&data))
	fmt.Println("part_2", part_2(&data))
}

type ScratchCard struct {
	cardID         int
	winningNumbers []int
	yourNumbers    []int
}

func parseLine(line *string) ScratchCard {

	winningNumbers := []int{}
	yourNumbers := []int{}
	cardID__numbers := strings.Split(*line, ": ")
	gameID, _ := strconv.Atoi(strings.TrimSpace(strings.SplitAfter(cardID__numbers[0], "d ")[1]))

	winning__yours := strings.Split(cardID__numbers[1], " | ")

	for index := range winning__yours[0] {

		//index of start of number is always a 3 multiple n(%3 == 0)
		if index%3 != 0 {
			continue
		}
		elem := winning__yours[0][index]
		next := winning__yours[0][index+1]
		chars := []byte{elem, next}
		num, _ := strconv.Atoi(strings.TrimSpace(string(chars)))

		winningNumbers = append(winningNumbers, num)
	}

	for index := range winning__yours[1] {

		//index of start of number is always a 3 multiple n(%3 == 0)
		if index%3 != 0 {
			continue
		}
		elem := winning__yours[1][index]
		next := winning__yours[1][index+1]
		chars := []byte{elem, next}
		num, _ := strconv.Atoi(strings.TrimSpace(string(chars)))

		yourNumbers = append(yourNumbers, num)
	}

	return ScratchCard{cardID: gameID, winningNumbers: winningNumbers, yourNumbers: yourNumbers}
}

func parse(input *string) []ScratchCard {
	lines := strings.Split(*input, "\n")
	return util.Map(lines, func(line string) ScratchCard {
		return parseLine(&line)
	})
}

func part_1(cards *[]ScratchCard) int {
	sum := 0
	for _, card := range *cards {

		nMatches := getCardMatches(card)
		if nMatches == 0 {
			continue
		}
		sum += IntPow(2, nMatches-1)
	}
	return sum
}

func part_2(cards *[]ScratchCard) int {
	values := []int{}
	i := 0
	last := len(*cards) - 1
	for {
		if last-i < 0 {
			break
		}
		values = append(values, getCardPoints(cards, &values, i, last))
		i++
	}

	//sum all values
	sum := 0
	for _, value := range values {
		sum += value
	}

	return sum
}

func getCardPoints(cards *[]ScratchCard, values *[]int, index int, last int) int {
	points := 1
	nMatches := getCardMatches((*cards)[last-index])

	for i := 1; i <= nMatches; i++ {
		points += (*values)[index-i]
	}
	return points
}

func getCardMatches(card ScratchCard) int {
	score := 0
	for _, wn := range card.winningNumbers {
		for _, yn := range card.yourNumbers {
			if wn == yn {
				score++
			}
		}
	}
	return score
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
