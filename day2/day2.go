package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var safeReports int
	var safeReportsWithDampener int
	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Fields(line)
		reportSafe := isReportSafe(report)

		if reportSafe {
			safeReports++
		} else if isReportSafeWithDampener(report) {
			safeReportsWithDampener++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("safe reports: ", safeReports)
	fmt.Println("safe reports with dampener: ", safeReports+safeReportsWithDampener)
}

func isSafeDistance(val1 string, val2 string) bool {
	int1, _ := strconv.Atoi(val1)
	int2, _ := strconv.Atoi(val2)
	diff := int1 - int2

	if diff == 0 {
		return false
	}

	if diff < 0 {
		diff = -diff
	}

	if 1 <= diff && diff <= 3 {
		return true
	} else {
		return false
	}
}

func isReportSafe(report []string) bool {
	decreased := false
	increased := false

	for i := 0; i < len(report)-1; i++ {
		val1, _ := strconv.Atoi(report[i])
		val2, _ := strconv.Atoi(report[i+1])
		diff := val1 - val2

		if diff < 0 {
			decreased = true
		}

		if 0 < diff {
			increased = true
		}

		safeDistance := isSafeDistance(report[i], report[i+1])

		if (increased && decreased) || !safeDistance {
			return false
		}
	}
	return true
}

func isReportSafeWithDampener(report []string) bool {
	for i := 0; i < len(report); i++ {
		reportClone := slices.Clone(report)
		modifiedReport := slices.Delete(reportClone, i, i+1)
		reportSafe := isReportSafe(modifiedReport)

		if reportSafe {
			return true
		}

	}

	return false
}
