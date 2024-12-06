package main

import (
	"fmt"
	"regexp"

	"github.com/wshivers93/AOC_2024/util"
)

func main() {
	fmt.Println("Hello day4")
	scanner, file := util.GetFileScanner(util.GetCurrentDirectory() + "/day4_input.txt")
	defer file.Close()

	count := 0
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		count += countMatchInLine(line)
	}
	count += countMatchAcrossLines(lines)

	fmt.Println("Result: ", count)
}

func countMatchInLine(line string) int {
	// separate regex into multiple statements to account for overlapping matches
	reForward := regexp.MustCompile(`(?:XMAS)`)
	reBackward := regexp.MustCompile(`(?:SAMX)`)
	forward := reForward.FindAll([]byte(line), -1)
	backward := reBackward.FindAll([]byte(line), -1)

	return len(forward) + len(backward)
}

func countMatchAcrossLines(lines []string) int {
	shouldLookUp := false
	shouldLookDown := true
	shouldLookLeft := false
	shouldLookRight := true
	count := 0
	fmt.Println(shouldLookLeft, shouldLookRight)
	fmt.Println(len(lines))

	for lineIndex, line := range lines {
		if lineIndex >= 3 {
			shouldLookUp = true
		}
		if lineIndex >= len(lines)-3 {
			shouldLookDown = false
		}

		for valIndex, val := range line {
			if valIndex < 3 {
				shouldLookLeft = false
			} else {
				shouldLookLeft = true
			}

			if valIndex >= len(line)-3 {
				shouldLookRight = false
			} else {
				shouldLookRight = true
			}

			// down
			if shouldLookDown && string(val) == "X" && string(lines[lineIndex+1][valIndex]) == "M" && string(lines[lineIndex+2][valIndex]) == "A" && string(lines[lineIndex+3][valIndex]) == "S" {
				count += 1
			}

			// up
			if shouldLookUp && string(val) == "X" && string(lines[lineIndex-1][valIndex]) == "M" && string(lines[lineIndex-2][valIndex]) == "A" && string(lines[lineIndex-3][valIndex]) == "S" {
				count += 1
			}

			// diagonal left
			if shouldLookLeft {
				// down
				if shouldLookDown && string(val) == "X" && string(lines[lineIndex+1][valIndex-1]) == "M" && string(lines[lineIndex+2][valIndex-2]) == "A" && string(lines[lineIndex+3][valIndex-3]) == "S" {
					count += 1
				}

				// up
				if shouldLookUp && string(val) == "X" && string(lines[lineIndex-1][valIndex-1]) == "M" && string(lines[lineIndex-2][valIndex-2]) == "A" && string(lines[lineIndex-3][valIndex-3]) == "S" {
					count += 1
				}
			}

			// diagonal right
			if shouldLookRight {
				// down
				if shouldLookDown && string(val) == "X" && string(lines[lineIndex+1][valIndex+1]) == "M" && string(lines[lineIndex+2][valIndex+2]) == "A" && string(lines[lineIndex+3][valIndex+3]) == "S" {
					count += 1
				}

				// up
				if shouldLookUp && string(val) == "X" && string(lines[lineIndex-1][valIndex+1]) == "M" && string(lines[lineIndex-2][valIndex+2]) == "A" && string(lines[lineIndex-3][valIndex+3]) == "S" {
					count += 1
				}
			}
		}
	}

	return count
}
