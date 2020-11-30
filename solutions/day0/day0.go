package main

import (
	"bufio"
	"fmt"
	"os"

	"../../shared"
)

func readFileContents(filePath string) []string {
	file, err := os.Open(filePath)

	if err == nil {
		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		return lines
	}

	panic(err)
}

func fuelRequired(mass int) int {
	// Go does floor division for ints
	return mass/3 - 2
}

func fuelRequiredPartTwo(mass int) int {
	fuelRequired := fuelRequired(mass)
	if fuelRequired > 0 {
		return fuelRequired + fuelRequiredPartTwo(fuelRequired)
	}
	return 0
}

func main() {

	// Solution to Day1 of 2019 to test everything is working
	floatLines := shared.ReadIntLines(0)

	totalFuelPartOne := 0

	for _, s := range floatLines {
		totalFuelPartOne += fuelRequired(s)
	}

	fmt.Println("Question 1\nTotal fuel required: ", int(totalFuelPartOne))

	totalFuelPartTwo := 0

	for _, s := range floatLines {
		totalFuelPartTwo += fuelRequiredPartTwo(s)
	}

	fmt.Println("Question 2\nTotal fuel required: ", totalFuelPartTwo)
}
