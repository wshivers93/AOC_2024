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
	file, err := os.Open("day1_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var firstList []int64
	var secondList []int64
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		firstInt, _ := strconv.ParseInt(split[0], 10, 64)
		secondInt, _ := strconv.ParseInt(split[1], 10, 64)
		firstList = append(firstList, firstInt)
		secondList = append(secondList, secondInt)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	slices.Sort(firstList)
	slices.Sort(secondList)

	totalDistance := getTotalDistance(firstList, secondList)
	fmt.Println("Total distance: ", totalDistance)
}

func getTotalDistance(arr1 []int64, arr2 []int64) int64 {
	var totalDistance int64
	entries := len(arr1)

	for i := 0; i < entries; i++ {
		distance := arr1[i] - arr2[i]

		if distance < 0 {
			distance = -distance
		}

		totalDistance += distance
	}

	return totalDistance
}
