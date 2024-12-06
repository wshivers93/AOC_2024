package util

import (
	"os"
	"fmt"
	"bufio"
)

func getFileScanner(fileName string) *bufio.Scanner {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	return scanner
}
