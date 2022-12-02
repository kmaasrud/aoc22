package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func a(file string) int {
    elves := strings.Split(file, "\n\n")

    maxcals := 1
    for _, foods := range elves {
	foods := strings.Split(foods, "\n")

	sum := 0
	for _, food := range foods {
	    cals, err := strconv.Atoi(food)

	    if err != nil {
		continue
	    }

	    sum += cals
	}

	if sum > maxcals {
	    maxcals = sum
	}
    }

    return maxcals
}

func b(file string) int {
    elves := strings.Split(file, "\n\n")
    var calsCarried []int
    for _, foods := range elves {
	foods := strings.Split(foods, "\n")

	sum := 0
	for _, food := range foods {
	    cals, err := strconv.Atoi(food)

	    if err != nil {
		continue
	    }

	    sum += cals
	}
	calsCarried = append(calsCarried, sum)
    }

    sort.Ints(calsCarried)

    sum := 0
    for _, cals := range calsCarried[len(calsCarried)-3:] {
	sum += cals
    }

    return sum
}

func main() {
    content, err := os.ReadFile("data/01.txt")

    if err != nil {
        log.Fatal(err)
    }

    maxcals := a(string(content))
    fmt.Printf("The elf carrying the most energy, carries %d calories.\n", maxcals)

    topThreeCals := b(string(content))
    fmt.Printf("The three elves carrying the most energy, carry %d calories\n", topThreeCals)
}
