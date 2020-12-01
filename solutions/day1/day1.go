package main

import (
	"errors"
	"fmt"
	"strconv"

	"../../shared"
)

const target = 2020

type pair struct {
	a int
	b int
}

func buildMap(lines []int) map[int]pair {
	length := len(lines)

	pairMap := make(map[int]pair)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			a := lines[i]
			b := lines[j]
			pairMap[a+b] = pair{a, b}
		}
	}
	return pairMap
}

// Day1Computer computes the solutions for day 1
type Day1Computer struct {
	input shared.Input
}

func (d *Day1Computer) part1() (shared.Result, error) {
	lines := shared.ToIntSlice(d.input)
	q1Map := make(map[int]bool)

	for _, l := range lines {
		toFind := target - l
		if q1Map[toFind] {
			return strconv.Itoa(toFind * l), nil
		}
		q1Map[l] = true
	}
	return "", errors.New("Could not find matching numbers")
}

func (d *Day1Computer) part2() (shared.Result, error) {
	lines := shared.ToIntSlice(d.input)
	q2Map := buildMap(lines)

	for _, l := range lines {
		if intPair, ok := q2Map[target-l]; ok {
			return strconv.Itoa(intPair.a * intPair.b * l), nil
		}
	}
	return "", errors.New("Could not find matching numbers")
}

func main() {

	fmt.Println("AoC 2020 Day 01")

	day1 := Day1Computer{shared.ReadStringLines(1)}

	ans1, err1 := day1.part1()
	fmt.Println("Question 1:", ans1, err1)

	ans2, err2 := day1.part2()
	fmt.Println("Question 1:", ans2, err2)
}
