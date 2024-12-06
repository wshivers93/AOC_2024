package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("day3_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	re := regexp.MustCompile(`(?:mul\()\d+\,\d+\)`)
	result := re.FindAll(file, -1)
	total := 0

	for _, val := range result {
		total += multiply(string(val))
	}
	fmt.Println("Part 1 result: ", total)

	// part 2
	re2 := regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	result2 := re2.FindAll(file, -1)
	insEnabled := true
	total2 := 0

	for _, val := range result2 {
		stringVal := string(val)
		if stringVal == "do()" {
			insEnabled = true
		}
		if stringVal == "don't()" {
			insEnabled = false
		}

		if insEnabled && strings.Contains(stringVal, "mul") {
			total2 += multiply(stringVal)
		}
	}

	fmt.Println("Part 2 result: ", total2)
}

func multiply(instruction string) int {
	re := regexp.MustCompile(`(?:mul\()\d+\,\d+\)`)
	ins := re.FindString(instruction)
	removeInstruction := strings.ReplaceAll(ins, "mul(", "")
	removeTrailing := strings.ReplaceAll(removeInstruction, ")", "")
	numbers := strings.Split(removeTrailing, ",")
	num1, _ := strconv.Atoi(numbers[0])
	num2, _ := strconv.Atoi(numbers[1])

	return num1 * num2
}
