package days

// Coordinate -
type Coordinate struct {
	x int
	y int
	z int
}

// GetNeighborsFromCoordinate -
func GetNeighborsFromCoordinate(coord Coordinate) []Coordinate {

	neighborPositions := make([]Coordinate, 0)

	for xDiff := -1; xDiff <= 1; xDiff++ {
		for yDiff := -1; yDiff <= 1; yDiff++ {
			for zDiff := -1; zDiff <= 1; zDiff++ {

				if !(xDiff == 0 && yDiff == 0 && zDiff == 0) {
					neighborPositions = append(neighborPositions, Coordinate{x: coord.x + xDiff, y: coord.y + yDiff, z: coord.z + zDiff})
				}

			}
		}
	}

	return neighborPositions
}
