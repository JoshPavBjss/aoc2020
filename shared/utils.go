package shared

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputFilePath = "/Users/josh.paveley/workspace/aoc/inputs/day%v.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInputPath(day int) string {
	return fmt.Sprintf(inputFilePath, day)
}

func openFileForDay(day int) *os.File {
	file, err := os.Open(getInputPath(day))
	check(err)
	return file
}

func getScannerForDay(day int) *bufio.Scanner {
	return bufio.NewScanner(openFileForDay(day))
}

// ReadLineFromFile - Reads the first line of the file for a given day
func ReadLineFromFile(day int) string {
	return getScannerForDay(day).Text()
}

// ReadStringLines - Reads all lines from the file to a string slice
func ReadStringLines(dayNumber int) []string {
	scanner := getScannerForDay(dayNumber)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// ReadIntLines - Reads all lines from the file to a int slice
func ReadIntLines(dayNumber int) []int {
	return ToIntSlice(ReadStringLines(dayNumber))
}

// ToIntSlice converts a slice of strings to a slice of ints
func ToIntSlice(stringSlice []string) []int {
	var lines []int
	for _, str := range stringSlice {
		converted, _ := strconv.Atoi(str)
		lines = append(lines, converted)
	}
	return lines
}
