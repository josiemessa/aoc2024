package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/josiemessa/aoc2024/utils"
)

func main() {
	lines := utils.ReadFileAsLines("input")
	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))
	for i, line := range lines {
		split := strings.Split(line, "   ")
		a1, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatalf("first entry line %d NaN\n", i)
		}
		list1[i] = a1

		a2, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("second entry line %d NaN\n", i)
		}
		list2[i] = a2
	}

	sort.Ints(list1)
	sort.Ints(list2)

	// DAY 1
	var result1 int

	for i, v := range list1 {
		diff := math.Abs(float64(v) - float64(list2[i]))
		result1 += int(diff)
	}

	fmt.Printf("Day 1, first part: %d\n", result1)

	// DAY 2
	sims := make(map[int]int)
	for _, v := range list2 {
		if _, ok := sims[v]; !ok {
			sims[v] = 1
		} else {
			sims[v]++
		}
	}

	var result2 int
	for _, v := range list1 {
		result2 += v * sims[v]
	}

	fmt.Printf("Day 2, second part: %d\n", result2)

}
