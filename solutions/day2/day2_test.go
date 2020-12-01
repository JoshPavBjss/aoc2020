package main

import (
	"testing"

	"../../shared"
)

func testInput() shared.Input {
	return shared.Input{}
}

func TestPart1(t *testing.T) {
	testDay := &Day2Computer{testInput()}

	res, err := testDay.part1()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "expected" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Day2Computer{testInput()}

	res, err := testDay.part2()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "expected" {
		t.Fatalf("Wrong result: %s", res)
	}
}
