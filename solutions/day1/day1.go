package main

import (
	"fmt"

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

func q1(lines []int) {
	q1Map := make(map[int]bool)
	for _, l := range lines {
		toFind := target - l
		if q1Map[toFind] {
			fmt.Printf("\nQuestion 1:\n%v * %v = %v\n", toFind, l, toFind*l)
			break
		}
		q1Map[l] = true
	}
}

func q2(lines []int) {
	q2Map := buildMap(lines)

	for _, l := range lines {
		if intPair, ok := q2Map[target-l]; ok {
			a := intPair.a
			b := intPair.b
			fmt.Printf("\nQuestion 2:\n%v * %v * %v = %v\n", a, b, l, a*b*l)
			break
		}
	}
}

func main() {

	fmt.Println("AoC 2020 Day 01")

	lines := shared.ReadIntLines(1)

	q1(lines)

	q2(lines)

}
