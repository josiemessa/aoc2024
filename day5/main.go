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
	rules := make(map[string][]rule, 0)

	var splitIndex int
	var result1 int
	var result2 int

	for i, line := range lines {
		if line == "" {
			splitIndex = i
			break
		}

		split := strings.Split(line, "|")
		if len(split) != 2 {
			log.Fatalf("Invalid pairs line %q\n", line)
		}

		if _, ok := rules[split[0]]; !ok {
			rules[split[0]] = []rule{rule{first: split[0], sec: split[1]}}
		} else {
			rules[split[0]] = append(rules[split[0]], rule{first: split[0], sec: split[1]})
		}

	}

	for _, line := range lines[splitIndex+1:] {
		pages := strings.Split(line, ",")

		allPagesOrdered := true

		brokenRules := make([]rule, 0)

		for _, p := range pages {
			matchingRules := rules[p]
			// see if this rule applies
			for _, rule := range matchingRules {
				if !slices.Contains(pages, rule.sec) {
					// rule does not apply
					continue
				}
				// check first comes before second
				i := slices.Index(pages, rule.first)
				if i == len(pages)-1 || !slices.Contains(pages[i+1:], rule.sec) {
					// second isn't after first, one of the rules is broken, stop looking
					allPagesOrdered = false
					brokenRules = append(brokenRules, rule)
				}
			}
		}

		if !allPagesOrdered {
			// try to order them
			for _, rule := range brokenRules {
				i := slices.Index(pages, rule.first)
				j := slices.Index(pages, rule.sec)
				if i > j { // rule may have been fixed by another swap
					pages[i] = rule.sec
					pages[j] = rule.first
				}
			}
			fmt.Print(line, " | ")
			for _, p := range pages {
				fmt.Print(p, ",")
			}
			fmt.Println()
		}

		// find the middle page
		mid := pages[int(len(pages)/2)]
		x, err := strconv.Atoi(mid)
		if err != nil {
			log.Fatalf("could not parse page number %q from line %q\n", mid, line)
		}

		if allPagesOrdered {
			result1 += x
		} else {
			result2 += x
		}
	}

	fmt.Println("Day 5 Part 1:", result1)
	fmt.Println("Day 5 Part 2:", result2)
}
