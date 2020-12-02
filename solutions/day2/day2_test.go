package main

import (
	"testing"

	"../../shared"
)

func testInput() shared.Input {
	return shared.Input{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
}

func TestPart1(t *testing.T) {
	testDay := &Day2Computer{testInput()}

	res, err := testDay.part1()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "2" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Day2Computer{testInput()}

	res, err := testDay.part2()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "1" {
		t.Fatalf("Wrong result: %s", res)
	}
}
