package main

import (
	"errors"
	"fmt"

	"../../shared"
)

// Day2Computer computes the solutions for day 1
type Day2Computer struct {
	input shared.Input
}

func (d *Day2Computer) part1() (shared.Result, error) {
	return "", errors.New("Not yet implemented")
}

func (d *Day2Computer) part2() (shared.Result, error) {
	return "", errors.New("Not yet implemented")
}

func main() {

	fmt.Println("AoC 2020 Day 01")

	day2 := Day2Computer{shared.ReadStringLines(2)}

	ans1, err1 := day2.part1()
	fmt.Println("Question 1:", ans1, err1)

	ans2, err2 := day2.part2()
	fmt.Println("Question 1:", ans2, err2)
}
