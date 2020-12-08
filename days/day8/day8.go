package days

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"../../shared"
)

// Day8Computer solves day8
type Day8Computer struct{}

type instructionMap map[int]instruction

type instruction struct {
	instructionType string
	value           int
}

func parseInstruction(line string) instruction {
	split := strings.Split(line, " ")
	parsedValue, _ := strconv.Atoi(split[1])
	return instruction{instructionType: split[0], value: parsedValue}
}

func buildInstructionMap(input shared.Input) instructionMap {
	instructions := make(instructionMap)
	for i, line := range input {
		instructions[i] = parseInstruction(line)
	}
	return instructions
}

func getValueAtLoopStart(instructions instructionMap) (int, error) {
	accumulator := 0
	i := 0

	visited := make(map[int]struct{})

	for i < len(instructions) {
		currentInstruction := instructions[i]
		if _, ok := visited[i]; ok {
			return accumulator, nil
		}
		visited[i] = struct{}{}

		if currentInstruction.instructionType == "acc" {
			accumulator += currentInstruction.value
		} else if currentInstruction.instructionType == "jmp" {
			i += currentInstruction.value
			continue
		}
		i++
	}
	return accumulator, errors.New("Should not have got here")
}

// Part1 of day 8
func (d *Day8Computer) Part1(input shared.Input) (shared.Result, error) {

	instructions := buildInstructionMap(input)

	fmt.Println(getValueAtLoopStart(instructions))

	return "", errors.New("Not yet implemented")
}

func createInstructionMapCopy(instructions instructionMap) instructionMap {
	newInstructions := make(instructionMap)
	for k, v := range instructions {
		newInstructions[k] = v
	}
	return newInstructions
}

func getSwitchedInstructionType(oldInstruction instruction) instruction {
	newInstructionType := "jmp"
	if oldInstruction.instructionType == "jmp" {
		newInstructionType = "nop"
	}
	return instruction{instructionType: newInstructionType, value: oldInstruction.value}
}

// Part2 of day 8
func (d *Day8Computer) Part2(input shared.Input) (shared.Result, error) {

	instructions := buildInstructionMap(input)

	for k, v := range instructions {
		if v.instructionType == "nop" || v.instructionType == "jmp" {
			newInstructions := createInstructionMapCopy(instructions)
			newInstructions[k] = getSwitchedInstructionType(v)

			val, err := getValueAtLoopStart(newInstructions)

			if err != nil {
				fmt.Println("This one worked, acc was:", val)
				break
			}
		}
	}

	return "", errors.New("Not yet implemented")
}
