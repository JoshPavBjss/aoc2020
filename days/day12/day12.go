package days

import (
	"fmt"
	"strconv"

	"../../shared"
)

// Day12Computer solves day12
type Day12Computer struct{}

type coords struct {
	x          int
	y          int
	currentDir int
}

// Part1 of day 10
func (d *Day12Computer) Part1(input shared.Input) (shared.Result, error) {

	pos := coords{0, 0, 1}

	for _, v := range input {

		action := rune(v[0])
		value, _ := strconv.Atoi(v[1:])

		applyAction(action, value, &pos)

		fmt.Println(v)
	}

	fmt.Println(pos)

	manhattenDistance := abs(pos.x) + abs(pos.y)

	return strconv.Itoa(manhattenDistance), nil
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

func applyAction(action rune, value int, pos *coords) {

	dirs := map[int]rune{
		0: 'N',
		1: 'E',
		2: 'S',
		3: 'W',
	}

	switch action {
	case 'N':
		pos.y += value
	case 'S':
		pos.y -= value
	case 'E':
		pos.x += value
	case 'W':
		pos.x -= value
	case 'L':
		pos.currentDir = ((pos.currentDir - value/90) + 4) % 4
	case 'R':
		pos.currentDir = ((pos.currentDir + value/90) + 4) % 4
	case 'F':
		applyAction(dirs[pos.currentDir], value, pos)
	}

}

// Part2 of day 10
func (d *Day12Computer) Part2(input shared.Input) (shared.Result, error) {

	return "", nil
}
