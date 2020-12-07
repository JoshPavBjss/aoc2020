package days

import (
	"regexp"
	"strconv"
	"strings"

	"../../shared"
)

// Day7Computer solves day7
type Day7Computer struct{}

type bagContent struct {
	amount int
	colour string
}

func getAmountAndColour(bag string) (int, string) {
	if bag == "no other bags." {
		return 0, ""
	}
	matches := regexp.MustCompile(`([0-9]{1,})\s([a-z].*?)\sbag`).FindAllStringSubmatch(bag, -1)
	amount, _ := strconv.Atoi(matches[0][1])
	return amount, matches[0][2]
}

func buildBagMap(input shared.Input) map[string][]bagContent {
	bags := make(map[string][]bagContent)

	for _, line := range input {

		bagContents := make([]bagContent, 0)

		keyValueSplit := strings.Split(line, "contain")

		key := strings.TrimSpace(strings.ReplaceAll(keyValueSplit[0], "bags", ""))

		values := strings.Split(keyValueSplit[1], ",")

		if strings.TrimSpace(keyValueSplit[1]) == "no other bags." {
			bagContents = nil
		} else {
			for _, subbag := range values {
				amount, colour := getAmountAndColour(strings.TrimSpace(subbag))
				bagContents = append(bagContents, bagContent{amount, colour})
			}
		}

		bags[key] = bagContents

	}
	return bags
}

func containsShinyGold(bags map[string][]bagContent, currentBag string) bool {
	bagContents := bags[currentBag]
	if bagContents == nil {
		return false
	}
	for _, bag := range bagContents {
		if bag.colour == "shiny gold" || containsShinyGold(bags, bag.colour) {
			return true
		}
	}
	return false
}

func calcBagsContained(bags map[string][]bagContent, currentBag string) int {
	bagContents := bags[currentBag]
	if bagContents == nil {
		return 0
	}
	bagCount := 0
	for _, bag := range bagContents {
		bagCount += bag.amount * (1 + calcBagsContained(bags, bag.colour))
	}
	return bagCount
}

// Part1 of day 7
func (d *Day7Computer) Part1(input shared.Input) (shared.Result, error) {

	bags := buildBagMap(input)

	shinyGold := 0

	for k := range bags {
		if containsShinyGold(bags, k) {
			shinyGold++
		}
	}

	return strconv.Itoa(shinyGold), nil
}

// Part2 of day 7
func (d *Day7Computer) Part2(input shared.Input) (shared.Result, error) {

	bags := buildBagMap(input)

	return strconv.Itoa(calcBagsContained(bags, "shiny gold")), nil
}
