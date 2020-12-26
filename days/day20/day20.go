package days

import (
	"fmt"
	"strconv"
	"strings"

	"../../shared"
)

// Day20Computer solves day20
type Day20Computer struct{}

func getTileID(line string) int {
	asInt, _ := strconv.Atoi(strings.ReplaceAll(line[0:len(line)-1], "Tile ", ""))
	return asInt
}

func createTilesFromInput(input shared.Input) []Tile {

	tiles := make([]Tile, 0)

	var newTile Tile

	for i, line := range input {

		if strings.Contains(line, "Tile") {
			newTile = Tile{id: getTileID(line)}
		} else if line == "" {
			tiles = append(tiles, newTile)
		} else {
			newTile.addImageData(line)

			if i == len(input)-1 {
				tiles = append(tiles, newTile)
			}
		}
	}

	return tiles
}

func getCorners(tiles []Tile) []int {

	cornerTiles := make([]int, 0)

	for _, tile := range tiles {
		count := 0
		for _, otherTile := range tiles {

			if !tile.equals(otherTile) {

				if tile.getNumberOfMachingCombinations(otherTile) > 1 {
					count++
				}

			}
		}
		if count == 2 {
			cornerTiles = append(cornerTiles, tile.id)
		}
	}
	return cornerTiles
}

// Part1 of day 20
func (d *Day20Computer) Part1(input shared.Input) (shared.Result, error) {

	tiles := createTilesFromInput(input)

	ans := 1

	for _, id := range getCorners(tiles) {
		ans *= id
	}

	return fmt.Sprintf("%v", ans), nil
}

// Part2 of day 20
func (d *Day20Computer) Part2(input shared.Input) (shared.Result, error) {

	return "", nil
}
