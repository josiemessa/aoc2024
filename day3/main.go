package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/josiemessa/aoc2024/utils"
)

const mulPrefix = "mul("
const doIntsr = "do()"
const dontInstr = "don't()"
const maxInstrLen = 12 // mul(123,123) (numbers between 1-3 digits, comma separated, no spaces)
const minInstrLen = 8  // mul(1,1)

func main() {
	input := utils.ReadFile("input")

	var result1 int
	do := true

	instructions := []string{"mul(", "do()", "don't()"}

	for i := 0; i < len(input); i++ {
		if i > len(input)-minInstrLen-1 {
			break
		}
		var foundInstruction string
		for _, inst := range instructions {
			x := string(input[i : len(inst)+i])
			if x == inst {
				foundInstruction = x
				break
			}
		}

		if foundInstruction != "" {
			if foundInstruction == doIntsr {
				do = true
				i += len(doIntsr) - 1
			}
			if foundInstruction == dontInstr {
				do = false
				i += len(dontInstr) - 1
			}

			if foundInstruction == mulPrefix && do {
				// pass in the maximal string length that could be a valid instruction depending on how much of the input is left
				l := math.Min(float64(len(input)-i), maxInstrLen)

				// make sure we copy slices because I don't want any slice fuckery to interfere - we have enough corrupted memory in this puzzle already
				candidate := make([]byte, int(l))
				copy(candidate, input[i:int(l)+i])

				res, parsedChars := parseMulInstruction(string(candidate))
				if res != -1 {
					result1 += res
					i += parsedChars - 1
				} else {
					// don't rescan the prefix we've already checked (the rest may be valid)
					i += len(mulPrefix)
				}
			}
		}
	}

	fmt.Println("Day 3 Part 1:", result1)
}

func parseMulInstruction(s string) (int, int) {
	// remove prefix as we already know it's OK
	s = s[len(mulPrefix):]

	// 3 is the earliest the closing bracket could be
	// 1,1)
	var foundClosingBracket bool
	for i := 3; i < len(s); i++ {
		if s[i] == ')' {
			// trim off everything after the closing bracket and the closing bracket
			s = s[:i]
			foundClosingBracket = true
			break
		}
	}
	if !foundClosingBracket {
		return -1, 0
	}

	// split on comma
	split := strings.Split(s, ",")
	if len(split) != 2 {
		// not valid
		return -1, 0
	}
	a, err := strconv.Atoi(split[0])
	if err != nil {
		return -1, 0
	}

	b, err := strconv.Atoi(split[1])
	if err != nil {
		return -1, 0
	}

	fmt.Printf("mul(%d*%d)\n", a, b)
	return a * b, len(mulPrefix) + len(s) + 1 // add the closing bracket back in
}
