package days

import (
	"fmt"

	"../../shared"
)

// Day24Computer solves day24
type Day24Computer struct{}

var startPos = coord{0, 0}

type coord struct {
	x int
	y int
}

func getDirections(line string) []string {

	directions := make([]string, 0)

	for i := 0; i < len(line); i++ {

		direction := string(line[i])

		if direction == "s" || direction == "n" {
			direction = line[i : i+2]
			i++
		}
		directions = append(directions, direction)

	}

	return directions
}

func getCoordsFromDirections(directions []string) coord {

	tilePos := startPos

	for _, direction := range directions {

		switch direction {
		case "ne":
			tilePos.x++
			tilePos.y++
		case "nw":
			tilePos.x--
			tilePos.y++
		case "se":
			tilePos.x++
			tilePos.y--
		case "sw":
			tilePos.x--
			tilePos.y--
		case "w":
			tilePos.x -= 2
		case "e":
			tilePos.x += 2
		}

	}

	return tilePos

}

func identifyTile(directions []string, floorMap *map[coord]bool) {

	tilePos := getCoordsFromDirections(directions)

	if tileFlipped, tileSet := (*floorMap)[tilePos]; tileSet {
		(*floorMap)[tilePos] = !tileFlipped
	} else {
		(*floorMap)[tilePos] = true
	}
}

// Part1 of day 24
func (d *Day24Computer) Part1(input shared.Input) (shared.Result, error) {

	floorMap := make(map[coord]bool)

	for _, line := range input {
		identifyTile(getDirections(line), &floorMap)
	}

	count := 0

	for _, v := range floorMap {
		if v {
			count++
		}
	}

	return fmt.Sprintf("%v", count), nil
}

// Part2 of day 24
func (d *Day24Computer) Part2(input shared.Input) (shared.Result, error) {

	return "", nil
}
