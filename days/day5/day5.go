package days

import (
	"errors"
	"math"
	"strconv"

	"../../shared"
)

const rowCharacters = 7

// Day5Computer solves day5
type Day5Computer struct{}

type max = int
type seatID = int

// SeatPosition stores the row and column of the seat
type SeatPosition struct {
	row int
	col int
}

// SeatMap stores all seat positions with their id
type SeatMap struct {
	seatMap map[seatID]SeatPosition
}

func (m SeatMap) idPresent(id seatID) bool {
	_, ok := m.seatMap[id]
	return ok
}

func (m SeatMap) idsPresent(ids []seatID) bool {
	for _, id := range ids {
		if !m.idPresent(id) {
			return false
		}
	}
	return true
}

func calculateRowNumber(rowPartition string) int {
	return binarySearch(rowPartition, 0, 127, "F", "B")
}

func calculateColumnNumber(columnPartition string) int {
	return binarySearch(columnPartition, 0, 7, "L", "R")
}

func binarySearch(input string, min int, max int, lower string, upper string) int {

	for _, l := range input {
		if string(l) == lower {
			max = ((max + min) / 2)
		} else if string(l) == upper {
			min = ((max + min) / 2)
		}
	}
	return max
}

func buildSeatMap(input shared.Input) (SeatMap, max) {

	SeatPositionMap := make(map[seatID]SeatPosition)
	maxSeatID := -1

	for _, line := range input {

		seatPosition := SeatPosition{
			row: calculateRowNumber(line[:7]),
			col: calculateColumnNumber(line[7:]),
		}

		id := getSeatID(seatPosition)

		maxSeatID = int(math.Max(float64(maxSeatID), float64(id)))

		SeatPositionMap[id] = seatPosition
	}

	return SeatMap{SeatPositionMap}, max(maxSeatID)
}

func getSeatID(seat SeatPosition) seatID {
	return seat.row*8 + seat.col
}

func inFirstRow(id int) bool {
	return id <= getSeatID(SeatPosition{0, 7})
}

func inLastRow(id int) bool {
	return id >= getSeatID(SeatPosition{127, 0})
}

// Part1 of day 5
func (d *Day5Computer) Part1(input shared.Input) (shared.Result, error) {
	_, max := buildSeatMap(input)
	return strconv.Itoa(max), nil
}

// Part2 of day 5
func (d *Day5Computer) Part2(input shared.Input) (shared.Result, error) {

	seatPositionMap, max := buildSeatMap(input)

	for id := 0; id < max; id++ {
		if !seatPositionMap.idPresent(id) && !inFirstRow(id) && !inLastRow(id) && seatPositionMap.idsPresent([]int{id - 1, id + 1}) {
			return strconv.Itoa(id), nil
		}
	}

	return "", errors.New("Could not find matching id")
}
