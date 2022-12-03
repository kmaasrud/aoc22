package main

import (
	"fmt"
	"os"
	"strings"
)

func priority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a') + 1
	} else if item >= 'A' && item <= 'Z' {
		return int(item - 'A') + 27
	}

	return 0
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		compartmentSize := len(line) / 2
		first := line[0:compartmentSize]
		second := line[compartmentSize:]

		Loop:
		for _, i := range first {
			for _, j := range second {
				if i == j {
					sum += priority(i)
					break Loop
				}
			}
		}
	}

	return sum
}

func part2(lines []string) int {
	sum := 0
	for i := 0; i + 3 < len(lines); i += 3 {
		first := lines[i]
		second := lines[i + 1]
		third := lines[i + 2]

		Loop:
		for _, f := range first {
			for _, s := range second {
				for _, t := range third {
					if f == s && s == t {
						sum += priority(f)
						break Loop
					}
				}
			}
		}
	}

	return sum
}

func main() {
	bytes, _ := os.ReadFile("data/03.txt")
	content := string(bytes)
	lines := strings.Split(content, "\n")

	fmt.Printf("Sum of priorities is %d\n", part1(lines))
	fmt.Printf("Sum of badges is %d\n", part2(lines))
}
