package main

//advent of Go-de 2023 - day 2

/*
	Ruleset:
	we have a bag filld with a random amount of
	"blue", "red" and "green" cubes

	we play a set of "games"
	each game is made of 1 or more "sets"

	in each set, the elf pulls out a random amout of cubes,
	and shows them to us

	Example imput:
	Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue

	each game is identified by a numeric "gameID"


	Goal:
	we need to determine which of all the games whould have been *pollible*
	if the bag only contained: 12 red cubes, 13 green cubes, and 14 blue cubes

	and sum together the ID's of those games
*/

import (
	util "advent-of-code-2023"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GameSet struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	GameID int
	Sets   []GameSet
}

func parse(input string) []Game {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	return util.Map(lines, parse_line)
}

func parse_line(line string) Game {
	game__sets := strings.Split(line, ":")
	//game id ok :D
	gameID, _ := strconv.Atoi(strings.Split(game__sets[0], " ")[1])
	sets := util.Map(strings.Split(game__sets[1], ";"), parse_set)
	return Game{Sets: sets, GameID: gameID}
}

func parse_set(set string) GameSet {
	var red int
	var green int
	var blue int

	for _, element := range strings.Split(set, ", ") {
		value__color := strings.Split(strings.TrimSpace(element), " ")

		switch value__color[1] {
		case "red":
			red, _ = strconv.Atoi(value__color[0])
		case "green":
			green, _ = strconv.Atoi(value__color[0])
		case "blue":
			blue, _ = strconv.Atoi(value__color[0])
		default:
			fmt.Println("color value:", value__color[1])
			panic("something strange happened here OwO")
		}
	}

	return GameSet{Red: red, Green: green, Blue: blue}
}

func part_1(games []Game) int {
	//if the bag only contained: 12 red cubes, 13 green cubes, and 14 blue cubes
	limit := GameSet{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	sum := 0

	for _, game := range games {
		if util.Every(game.Sets, func(set GameSet) bool {
			return set.Red <= limit.Red && set.Green <= limit.Green && set.Blue <= limit.Blue
		}) {
			sum += game.GameID
		}
	}

	return sum
}

/*
Part 2 Goals:
for each game, find the mininum amout of cubes for each color to make it possible,
then get the "power" of the resulting set. n.red*n.green*n.blue

the final goal is to add thogether the "resulting power" of all the games.
*/
func part_2(games []Game) int {
	sum := 0

	for _, game := range games {
		minRed, minGreen, minBlue := min_blocks(game.Sets)
		pow := minRed * minGreen * minBlue
		sum += pow
	}

	return sum
}

func min_blocks(sets []GameSet) (int, int, int) {
	maxRed := sets[0].Red
	maxGreen := sets[0].Green
	maxBlue := sets[0].Blue

	for i := 0; i < len(sets); i++ {
		elem := sets[i]
		if elem.Red > maxRed {
			maxRed = elem.Red
		}
		if elem.Green > maxGreen {
			maxGreen = elem.Green
		}
		if elem.Blue > maxBlue {
			maxBlue = elem.Blue
		}
	}

	return maxRed, maxGreen, maxBlue
}

func main() {
	content, _ := os.ReadFile("day2/input")
	fmt.Println("part_1", part_1(parse(string(content))))
	fmt.Println("part_2", part_2(parse(string(content))))
}
