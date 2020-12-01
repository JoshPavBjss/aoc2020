package main

import (
	"fmt"
	"strconv"

	"../../shared"
)

// Day0Computer computes the solutions for day 0
type Day0Computer struct {
	input shared.Input
}

func (d *Day0Computer) part1() (shared.Result, error) {

	lines := shared.ToIntSlice(d.input)
	totalFuelPartOne := 0

	for _, s := range lines {
		totalFuelPartOne += fuelRequired(s)
	}
	return strconv.Itoa(totalFuelPartOne), nil
}

func (d *Day0Computer) part2() (shared.Result, error) {

	lines := shared.ToIntSlice(d.input)
	totalFuelPartTwo := 0

	for _, s := range lines {
		totalFuelPartTwo += fuelRequiredPartTwo(s)
	}
	return strconv.Itoa(totalFuelPartTwo), nil
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
	day0Computer := Day0Computer{shared.ReadStringLines(0)}

	pt1, _ := day0Computer.part1()

	fmt.Println("Question 1\nTotal fuel required: ", pt1)

	pt2, _ := day0Computer.part2()

	fmt.Println("Question 1\nTotal fuel required: ", pt2)
}
