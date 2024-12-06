package util

import (
	"os"
	"fmt"
	"bufio"
)

func GetFileScanner(filePath string) (bufio.Scanner, os.File) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	return *scanner, *file
}

func GetCurrentDirectory() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return pwd
}
