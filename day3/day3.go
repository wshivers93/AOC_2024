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
}

func multiply(instruction string) int {
	removeInstruction := strings.ReplaceAll(instruction, "mul(", "")
	removeTrailing := strings.ReplaceAll(removeInstruction, ")", "")
	numbers := strings.Split(removeTrailing, ",")
	num1, _ := strconv.Atoi(numbers[0])
	num2, _ := strconv.Atoi(numbers[1])

	return num1 * num2
}
