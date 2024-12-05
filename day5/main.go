package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/josiemessa/aoc2024/utils"
)

type rule struct {
	first string
	sec   string
}

func main() {
	lines := utils.ReadFileAsLines("input")
	rules := make([]rule, 0)

	var splitIndex int
	var result1 int

	for i, line := range lines {
		if line == "" {
			splitIndex = i
			break
		}

		split := strings.Split(line, "|")
		if len(split) != 2 {
			log.Fatalf("Invalid pairs line %q\n", line)
		}

		rules = append(rules, rule{first: split[0], sec: split[1]})
	}

	for _, line := range lines[splitIndex+1:] {
		pages := strings.Split(line, ",")

		allPagesOrdered := true

		for _, rule := range rules {
			// see if this rule applies
			if slices.Contains(pages, rule.first) && slices.Contains(pages, rule.sec) {
				// check first comes before second
				i := slices.Index(pages, rule.first)
				if i == len(pages)-1 || !slices.Contains(pages[i+1:], rule.sec) {
					// second isn't after first, one of the rules is broken, stop looking
					allPagesOrdered = false
					break
				}
			}
			// rule does not apply
		}

		if !allPagesOrdered {
			continue
		}

		// find the middle page
		mid := pages[int(len(pages)/2)]
		x, err := strconv.Atoi(mid)
		if err != nil {
			log.Fatalf("could not parse page number %q from line %q\n", mid, line)
		}
		fmt.Println(line)
		result1 += x
	}

	fmt.Println("Day 5 Part 1:", result1)
}
