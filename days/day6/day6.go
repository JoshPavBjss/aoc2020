package days

import (
	"strconv"

	"../../shared"
)

// Day6Computer solves day6
type Day6Computer struct{}

type set = map[rune]struct{}

// Part1 of day 6
func (d *Day6Computer) Part1(input shared.Input) (shared.Result, error) {

	answers := make(set)

	uniqueAnswers := 0

	for i, line := range input {

		if len(line) > 0 {
			for _, char := range line {
				answers[char] = struct{}{}
			}
		}

		if len(line) == 0 || i == len(input)-1 {
			uniqueAnswers += len(answers)
			answers = make(set)
			continue
		}
	}
	return strconv.Itoa(uniqueAnswers), nil
}

// Part2 of day 6
func (d *Day6Computer) Part2(input shared.Input) (shared.Result, error) {

	answers := make(set)

	uniqueAnswers := 0

	firstLine := true

	for i, line := range input {
		newAnswers := make(set)

		if len(line) > 0 {
			for _, char := range line {
				if firstLine {
					newAnswers[char] = struct{}{}
				} else if _, ok := answers[char]; ok {
					newAnswers[char] = struct{}{}
				}
			}
			answers = newAnswers
			firstLine = false
		}

		if len(line) == 0 || i == len(input)-1 {
			uniqueAnswers += len(answers)
			answers = make(map[rune]struct{})
			firstLine = true
			continue
		}
	}
	return strconv.Itoa(uniqueAnswers), nil
}
