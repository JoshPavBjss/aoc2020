package main

import (
	"testing"

	"../../shared"
)

func TestPart1_1(t *testing.T) {
	testDay := &Day0Computer{shared.Input{"12"}}

	res, err := testDay.part1()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "2" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_2(t *testing.T) {
	testDay := &Day0Computer{shared.Input{"14"}}

	res, err := testDay.part1()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "2" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_3(t *testing.T) {
	testDay := &Day0Computer{shared.Input{"1969"}}

	res, err := testDay.part1()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "654" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_4(t *testing.T) {
	testDay := &Day0Computer{shared.Input{"100756"}}

	res, err := testDay.part1()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "33583" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_1(t *testing.T) {
	testDay := &Day0Computer{shared.Input{"14"}}

	res, err := testDay.part2()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "2" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_2(t *testing.T) {
	testDay := &Day0Computer{shared.Input{"1969"}}

	res, err := testDay.part2()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "966" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_3(t *testing.T) {
	testDay := &Day0Computer{shared.Input{"100756"}}

	res, err := testDay.part2()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "50346" {
		t.Fatalf("Wrong result: %s", res)
	}
}
