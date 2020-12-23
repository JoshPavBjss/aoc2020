package days

import (
	"fmt"
	"strconv"
	"strings"

	"../../shared"
)

// Day19Computer solves day19
type Day19Computer struct{}

type rule struct {
	subRules       [][]int
	letter         string
	ruleMap        *map[int]rule
	cachedPatterns *map[int][]string
}

func (r *rule) getRulePatterns() []string {

	if r.subRules == nil {
		return []string{r.letter}
	}

	toReturn := make([]string, 0)

	for _, subRule := range r.subRules {

		patterns := make([]string, 0)

		for _, ruleNumber := range subRule {
			theRule := (*r.ruleMap)[ruleNumber]

			var rulePatterns []string
			var ok bool

			if rulePatterns, ok = (*r.cachedPatterns)[ruleNumber]; !ok {
				rulePatterns = theRule.getRulePatterns()
			}

			if len(patterns) == 0 {
				patterns = append(patterns, rulePatterns...)
			} else {

				updatedPatterns := make([]string, 0)

				for _, pattern := range patterns {
					for _, rulePattern := range rulePatterns {
						updatedPatterns = append(updatedPatterns, pattern+rulePattern)
					}
				}
				patterns = updatedPatterns
			}

		}

		toReturn = append(toReturn, patterns...)
	}

	return toReturn
}

func parseRules(input shared.Input) (map[int]rule, int) {

	rules := make(map[int]rule, 0)
	cachedPatterns := make(map[int][]string)

	for i, line := range input {

		if line == "" {
			return rules, i
		}

		splitOne := strings.Split(line, ":")

		ruleNumber, _ := strconv.Atoi(splitOne[0])

		if strings.Contains(splitOne[1], "\"") {

			letter := strings.Replace(splitOne[1], "\"", "", 2)

			rules[ruleNumber] = rule{subRules: nil, letter: strings.TrimSpace(letter), ruleMap: &rules, cachedPatterns: &cachedPatterns}

		} else {
			subRuleString := strings.Split(splitOne[1], "|")

			subRules := make([][]int, 0)

			for _, ruleSet := range subRuleString {

				subRuleNumbers := make([]int, 0)

				for _, subRuleNumber := range strings.Split(strings.TrimSpace(ruleSet), " ") {
					asNum, _ := strconv.Atoi(subRuleNumber)
					subRuleNumbers = append(subRuleNumbers, asNum)
				}

				subRules = append(subRules, subRuleNumbers)
			}

			rules[ruleNumber] = rule{subRules: subRules, letter: "", ruleMap: &rules, cachedPatterns: &cachedPatterns}

		}

	}
	return nil, 1
}

func matchesRule0(ruleMap map[int]rule, message string) bool {

	theRule := ruleMap[0]

	for _, rulePattern := range theRule.getRulePatterns() {
		if rulePattern == message {
			return true
		}
	}

	return false
}

// Part1 of day 19
func (d *Day19Computer) Part1(input shared.Input) (shared.Result, error) {

	rules, i := parseRules(input)

	matchingMessages := 0

	for i = i + 1; i < len(input); i++ {
		if matchesRule0(rules, input[i]) {
			matchingMessages++
		}
	}

	return fmt.Sprintf("%v", matchingMessages), nil
}

// Part2 of day 19
func (d *Day19Computer) Part2(input shared.Input) (shared.Result, error) {

	return "", nil
}
