package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../shared"
)

// Day2Computer computes the solutions for day 1
type Day2Computer struct {
	input shared.Input
}

// PasswordPolicy policy stores the information used to validate a password
type PasswordPolicy struct {
	min    int
	max    int
	letter string
}

func getPolicy(line string) (PasswordPolicy, string) {
	split := strings.Split(line, ":")
	lhs := strings.Split(split[0], " ")
	minMax := strings.Split(lhs[0], "-")

	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	pw := PasswordPolicy{min, max, lhs[1]}

	return pw, strings.TrimSpace(split[1])
}

type validator func(policy PasswordPolicy, pw string) bool

func validatePassword1(policy PasswordPolicy, pw string) bool {
	occurrences := strings.Count(pw, policy.letter)
	return occurrences >= policy.min && occurrences <= policy.max
}

func validatePasswordPt2(policy PasswordPolicy, pw string) bool {
	minIndexMatches := string(pw[policy.min-1]) == policy.letter
	maxIndexMatches := string(pw[policy.max-1]) == policy.letter
	return (minIndexMatches && !maxIndexMatches) || (!minIndexMatches && maxIndexMatches)
}

func validate(toValidate shared.Input, fn validator) string {
	validPasswords := 0

	for _, line := range toValidate {

		if fn(getPolicy(line)) {
			validPasswords++
		}
	}
	return strconv.Itoa(validPasswords)
}

func (d *Day2Computer) part1() (shared.Result, error) {
	return validate(d.input, validatePassword1), nil
}

func (d *Day2Computer) part2() (shared.Result, error) {
	return validate(d.input, validatePasswordPt2), nil
}

func main() {

	fmt.Println("AoC 2020 Day 01")

	day2 := Day2Computer{shared.ReadStringLines(2)}

	ans1, err1 := day2.part1()
	fmt.Println("Question 1:", ans1, err1)

	ans2, err2 := day2.part2()
	fmt.Println("Question 2:", ans2, err2)
}
