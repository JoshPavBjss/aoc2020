package main

import (
	"fmt"
	"time"

	today "../days/day8"
	shared "../shared"
)

func runForDay(day shared.Day, input shared.Input) {

	start := time.Now()

	ans1, err1 := day.Part1(input)
	fmt.Println("Question 1:", ans1, err1)

	ans2, err2 := day.Part2(input)
	fmt.Println("Question 2:", ans2, err2)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}

func main() {

	day := 8

	fmt.Println("AoC 2020 Day", day)

	runForDay(&today.Day8Computer{}, shared.Input(shared.ReadStringLines(day)))
}
