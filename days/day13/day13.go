package days

import (
	"errors"
	"strconv"
	"strings"

	"../../shared"
)

// Day13Computer solves day13
type Day13Computer struct{}

// Part1 of day 10
func (d *Day13Computer) Part1(input shared.Input) (shared.Result, error) {

	earliestDepartTime, _ := strconv.Atoi(input[0])
	busesInService := shared.ToIntSlice(strings.Split(strings.Replace(input[1], ",x", "", -1), ","))

	currentTime := earliestDepartTime

	for true {

		for _, bus := range busesInService {
			if currentTime%bus == 0 {
				return strconv.Itoa((currentTime - earliestDepartTime) * bus), nil
			}
		}
		currentTime++
	}

	return "", errors.New("Could not find available bus")
}

type busDepart struct {
	busNo  int
	offset int
}

func earliestBusDepart(busTimes []busDepart, start, increment int) int {

	timeStamp := start

	for true {
	busLoop:
		for i, bus := range busTimes {
			if (timeStamp+bus.offset)%bus.busNo != 0 {
				break busLoop
			}
			if i == len(busTimes)-1 {
				return timeStamp
			}
		}
		timeStamp += increment
	}
	return -1
}

func getIncrement(buses []busDepart) int {
	increment := 1
	for _, bus := range buses {
		increment *= bus.busNo
	}
	return increment
}

// Part2 of day 10
func (d *Day13Computer) Part2(input shared.Input) (shared.Result, error) {

	buses := strings.Split(input[1], ",")

	busDeparts := make([]busDepart, 0)

	for i, bus := range buses {
		if busNo, err := strconv.Atoi(bus); err == nil {
			busDeparts = append(busDeparts, busDepart{busNo: busNo, offset: i})
		}
	}

	startTime := busDeparts[0].busNo

	for i := range busDeparts {

		startTime = earliestBusDepart(busDeparts[0:i+1], startTime, getIncrement(busDeparts[0:i]))
	}

	return strconv.Itoa(startTime), nil
}
