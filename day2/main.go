package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/josiemessa/aoc2024/utils"
)

func main() {
	var result1 int
	var result2 int

	lines := utils.ReadFileAsLines("input")
	for _, line := range lines {
		split := strings.Split(line, " ")
		levels := utils.SliceAtoi(split)

		increasing := levels[0] < levels[1]
		var failed bool

		for i := 0; i < len(levels)-1; i++ {
			failed = testLevel(levels[i], levels[i+1], increasing)
			if failed {
				break
			}
		}

		if !failed {
			result1++
			result2++
			println(line)
			continue
		}

		var failedAgain bool
		for i := 0; i <= len(levels)-1; i++ {
			// remove indices one at a time
			newLevels := make([]int, i)
			copy(newLevels, levels[:i])
			newLevels = append(newLevels, levels[i+1:]...)
			increasing = newLevels[0] < newLevels[1]
			for j := 0; j < len(newLevels)-1; j++ {
				failedAgain = testLevel(newLevels[j], newLevels[j+1], increasing)
				if failedAgain {
					break
				}
			}
			if !failedAgain {
				result2++
				println(line, "failed index:", i)
				break
			}
		}
	}

	fmt.Println("Day 2 Part 1:", result1)
	fmt.Println("Day 2 Part 2:", result2)

}

func testLevel(a, b int, increasing bool) bool {
	diff := a - b
	if math.Abs(float64(diff)) == 0 || math.Abs(float64(diff)) > 3 {
		return true
	}
	if a > b && increasing {
		return true
	}
	if b > a && !increasing {
		return true
	}
	return false
}
