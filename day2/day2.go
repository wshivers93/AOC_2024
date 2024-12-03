package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("day2_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var safeReports int
	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Fields(line)
		shouldBeDecreasing := isDecreasing(report[0], report[1])
		isSafe := true

		for i := 0; i < len(report) - 1; i++ {
			isDec := isDecreasing(report[i], report[i+1])
			isSafeDist := isSafeDistance(report[i], report[i+1])

			if (isDec != shouldBeDecreasing || !isSafeDist) {
				isSafe = false
				break
			}
		}

		if (isSafe) {
			safeReports++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("safe reports: ", safeReports)
}

func isDecreasing(val1 string, val2 string) bool {
	int1, _ := strconv.Atoi(val1) 
	int2, _ := strconv.Atoi(val2) 
	diff := int1 - int2 

	if (diff < 0) {
		return false 
	}

	return true 
}

func isSafeDistance(val1 string, val2 string) bool {
	int1, _ := strconv.Atoi(val1) 
	int2, _ := strconv.Atoi(val2) 
	diff := int1 - int2 

	if (diff == 0) {
		return false
	}

	if (diff < 0) {
		diff = -diff
	}

	if (1 <= diff && diff <= 3) {
		return true
	} else {
		return false
	}
}
