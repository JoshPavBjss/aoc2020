package days

import (
	"strconv"

	"../../shared"
)

const floor = '.'
const emptySeat = 'L'
const occupiedSeat = '#'

// Day11Computer solves day11
type Day11Computer struct{}

func getNewSeatState(seatLayout [][]rune, row, col int) rune {

	currentSeat := (seatLayout)[row][col]

	if currentSeat == emptySeat && adjacentSeatsOccupiedCount(seatLayout, row, col) == 0 {
		return occupiedSeat
	} else if currentSeat == occupiedSeat && adjacentSeatsOccupiedCount(seatLayout, row, col) >= 4 {
		return emptySeat
	}
	return currentSeat
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func adjacentSeatsOccupiedCount(seatLayout [][]rune, row, col int) int {

	count := 0

	minRow := max(0, row-1)
	maxRow := min(len(seatLayout)-1, row+1)

	minCol := max(0, col-1)
	maxCol := min(len(seatLayout[row])-1, col+1)

	for rowIt := minRow; rowIt <= maxRow; rowIt++ {
		for colIt := minCol; colIt <= maxCol; colIt++ {
			if rowIt != row || colIt != col {
				if seatLayout[rowIt][colIt] == occupiedSeat {
					count++
				}
			}
		}
	}

	return count
}

func createSeatLayout(input shared.Input) [][]rune {

	seatLayout := make([][]rune, len(input))

	for row, line := range input {
		seatLayout[row] = make([]rune, len(line))
		for col, char := range line {
			seatLayout[row][col] = char
		}
	}

	return seatLayout
}

func getUpdatedLayout(seatLayout [][]rune) [][]rune {

	newSeatLayout := make([][]rune, len(seatLayout))

	for rowIt, row := range seatLayout {
		newSeatLayout[rowIt] = make([]rune, len(row))
		for colIt := range row {
			newSeat := getNewSeatState(seatLayout, rowIt, colIt)
			newSeatLayout[rowIt][colIt] = newSeat
		}
	}

	return newSeatLayout

}

func sliceEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func seatLayoutsEqual(a, b [][]rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !sliceEqual(v, b[i]) {
			return false
		}
	}
	return true
}

func getOccupiedCount(seatLayout [][]rune) int {
	count := 0
	for _, row := range seatLayout {
		for _, seat := range row {
			if seat == occupiedSeat {
				count++
			}
		}
	}
	return count
}

// Part1 of day 10
func (d *Day11Computer) Part1(input shared.Input) (shared.Result, error) {

	seatLayout := createSeatLayout(input)

	for {
		newSeatLayout := getUpdatedLayout(seatLayout)

		if seatLayoutsEqual(seatLayout, newSeatLayout) {
			break
		}

		// Reset
		seatLayout = newSeatLayout
	}

	return strconv.Itoa(getOccupiedCount(seatLayout)), nil
}

// Part2 of day 10
func (d *Day11Computer) Part2(input shared.Input) (shared.Result, error) {

	return "", nil
}
