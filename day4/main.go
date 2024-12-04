package main

import (
	"fmt"

	"github.com/josiemessa/aoc2024/utils"
)

type direction struct {
	down  int
	right int
}

type puzzle struct {
	maxI   int
	maxJ   int
	lines  []string
	output [][]byte
}

func main() {
	Part1()
	Part2()
}

func Part1() {
	p := puzzle{
		lines: utils.ReadFileAsLines("test-input"),
	}
	p.maxI = len(p.lines) - 1
	p.maxJ = len(p.lines[0]) - 1
	p.output = make([][]byte, p.maxI+1)
	for i := range p.output {
		p.output[i] = make([]byte, p.maxJ+1)
	}

	var result1 int

	// i is the row, j is the column
	for i, line := range p.lines {
		for j := range line {
			if p.output[i][j] == 0 {
				p.output[i][j] = '.'
			}
			if line[j] != 'X' {
				continue
			}
			directions := p.lookAround(i, j, 'M')

			for _, d := range directions {
				if p.searchDirection(i+d.down, j+d.right, 'A', d) {
					if p.searchDirection(i+d.down+d.down, j+d.right+d.right, 'S', d) {
						p.output[i][j] = 'X'
						p.output[i+d.down][j+d.right] = 'M'
						p.output[i+d.down+d.down][j+d.right+d.right] = 'A'
						p.output[i+d.down+d.down+d.down][j+d.right+d.right+d.right] = 'S'
						result1++
					}
				}
			}
		}
	}

	fmt.Println("Day 4 Part 1:", result1)
	for _, v := range p.output {
		fmt.Println(string(v))
	}
}

func Part2() {
	p := puzzle{
		lines: utils.ReadFileAsLines("input"),
	}
	p.maxI = len(p.lines) - 1
	p.maxJ = len(p.lines[0]) - 1
	p.output = make([][]byte, p.maxI+1)
	for i := range p.output {
		p.output[i] = make([]byte, p.maxJ+1)
	}

	var result2 int

	// i is the row, j is the column
	for i, line := range p.lines {
		for j := range line {
			if p.output[i][j] == 0 {
				p.output[i][j] = '.'
			}
			if line[j] != 'A' {
				continue
			}
			directions := p.lookCorners(i, j, 'M')
			if len(directions) != 2 {
				continue
			}

			d0 := directions[0].invert()
			d1 := directions[1].invert()
			if p.searchDirection(i, j, 'S', d0) && p.searchDirection(i, j, 'S', d1) {
				result2++
			}

		}
	}

	fmt.Println("Day 4 Part 2:", result2)
	// for _, v := range p.output {
	// 	fmt.Println(string(v))
	// }
}

// the 'X' can be the start of multiple words, so return all directions
func (p puzzle) lookAround(i, j int, char byte) []direction {
	results := make([]direction, 0)

	searchDirections := []direction{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

	for _, d := range searchDirections {
		if p.searchDirection(i, j, char, d) {
			results = append(results, d)
		}
	}

	return results
}

// the 'X' can be the start of multiple words, so return all directions
func (p puzzle) lookCorners(i, j int, char byte) []direction {
	results := make([]direction, 0)

	searchDirections := []direction{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

	for _, d := range searchDirections {
		if p.searchDirection(i, j, char, d) {
			results = append(results, d)
		}
	}

	return results
}

func (p puzzle) searchDirection(i, j int, char byte, d direction) bool {
	newI := i + d.down
	newJ := j + d.right

	if newI >= 0 && newI <= p.maxI && newJ >= 0 && newJ <= p.maxJ {
		return p.lines[newI][newJ] == char
	}
	return false
}

func (d direction) invert() direction {
	return direction{d.down * -1, d.right * -1}
}
