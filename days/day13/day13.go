package days

import (
	"errors"
	"strconv"
	"strings"

	"../../shared"
)

// Day13Computer solves day13
type Day13Computer struct{}

type busSchedule struct {
	busNo  int
	offset int
}

func createBusSchedule(input shared.Input) []busSchedule {
	busSchedules := make([]busSchedule, 0)

	for i, bus := range strings.Split(input[1], ",") {
		if busNo, err := strconv.Atoi(bus); err == nil {
			busSchedules = append(busSchedules, busSchedule{busNo: busNo, offset: i})
		}
	}
	return busSchedules
}

func busDepartsAtTime(busNo, currentTime int) bool {
	return currentTime%busNo == 0
}

// Part1 of day 10
func (d *Day13Computer) Part1(input shared.Input) (shared.Result, error) {

	earliestDepartTime, _ := strconv.Atoi(input[0])

	busesInService := createBusSchedule(input)

	currentTime := earliestDepartTime

	for true {

		for _, bus := range busesInService {
			if busDepartsAtTime(bus.busNo, currentTime) {
				return strconv.Itoa((currentTime - earliestDepartTime) * bus.busNo), nil
			}
		}
		currentTime++
	}

	return "", errors.New("Could not find available bus")
}

func earliestBusDepart(busTimes []busSchedule, start, increment int) int {

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

func getIncrement(buses []busSchedule) int {
	increment := 1
	for _, bus := range buses {
		increment *= bus.busNo
	}
	return increment
}

// Part2 of day 10
func (d *Day13Computer) Part2(input shared.Input) (shared.Result, error) {

	buses := strings.Split(input[1], ",")

	busSchedules := make([]busSchedule, 0)

	for i, bus := range buses {
		if busNo, err := strconv.Atoi(bus); err == nil {
			busSchedules = append(busSchedules, busSchedule{busNo: busNo, offset: i})
		}
	}

	startTime := busSchedules[0].busNo

	for i := range busSchedules {

		startTime = earliestBusDepart(busSchedules[0:i+1], startTime, getIncrement(busSchedules[0:i]))
	}

	return strconv.Itoa(startTime), nil
}
