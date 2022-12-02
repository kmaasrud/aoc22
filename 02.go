package main

import (
	"fmt"
	"os"
	"strings"
)

var drawMap = map[string]string {
	"X": "A",
	"Y": "B",
	"Z": "C",
}

var winMap = map[string]string {
	"X": "C",
	"Y": "A",
	"Z": "B",
}

func game(a string, b string) int {
	if winMap[b] == a {
		return 6
	} else if drawMap[b] == a {
		return 3
	} else {
		return 0
	}
}

func shapeScore(shape string) int {
	switch shape {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}
	return 0
}

func score(a string, b string) int {
	return shapeScore(b) + game(a, b)
}

func part1(content string) int {
	lines := strings.Split(string(content), "\n")

	sum := 0
	for _, line := range lines {
		line := strings.Split(line, " ")

		if len(line) < 2 {
			continue
		}

		sum += score(line[0], line[1])
	}

	return sum
}

var draw = map[string]string {
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var win = map[string]string {
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var lose = map[string]string {
	"A": "Z",
	"B": "X",
	"C": "Y",
}

func part2(content string) int {
	lines := strings.Split(string(content), "\n")

	sum := 0
	for _, line := range lines {
		line := strings.Split(line, " ")

		if len(line) < 2 {
			continue
		}

		var choice string
		switch line[1] {
		case "X":
			choice = lose[line[0]]
		case "Y":
			choice = draw[line[0]]
		case "Z":
			choice = win[line[0]]
		}

		sum += score(line[0], choice)
	}

	return sum
}

func main() {
	bytes, _ := os.ReadFile("data/02.txt")
	content := string(bytes)

	score := part1(content)
	fmt.Printf("Part 1 score is %d\n", score)

	score = part2(content)
	fmt.Printf("Part 2 score is %d\n", score)
}
