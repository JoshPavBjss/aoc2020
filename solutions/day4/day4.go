package days

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"../../shared"
)

var passwordFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	"cid",
}

var requiredPasswordFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

// Day4Computer computes the solutions for day 4
type Day4Computer struct{}

type passportData struct {
	BYR string
	IYR string
	EYR string
	HGT string
	HCL string
	ECL string
	PID string
	CID string
}

func (p *passportData) hasRequiredFields() bool {
	return len(p.BYR) > 0 &&
		len(p.IYR) > 0 &&
		len(p.EYR) > 0 &&
		len(p.HGT) > 0 &&
		len(p.HCL) > 0 &&
		len(p.ECL) > 0 &&
		len(p.PID) > 0
}

func (p *passportData) isValid() bool {
	return p.hasRequiredFields() &&
		p.birthYearValid() &&
		p.issueYearValid() &&
		p.expYearValid() &&
		p.heightValid() &&
		p.hairColourValid() &&
		p.eyeColourValid() &&
		p.passportIDValid()
}

func (p *passportData) birthYearValid() bool {
	return validateInRange(p.BYR, 1920, 2002)
}

func (p *passportData) issueYearValid() bool {
	return validateInRange(p.IYR, 2010, 2020)
}

func (p *passportData) expYearValid() bool {
	return validateInRange(p.EYR, 2020, 2030)
}

func (p *passportData) heightValid() bool {

	if height, unit := parseHeight(p.HGT); unit == "in" {
		return validateInRange(height, 59, 76)
	} else if unit == "cm" {
		return validateInRange(height, 150, 193)
	}
	return false
}

func (p *passportData) hairColourValid() bool {
	return regexp.MustCompile(`#[a-f0-9]{6}`).MatchString(p.HCL)
}

func (p *passportData) eyeColourValid() bool {
	colours := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, colour := range colours {
		if colour == p.ECL {
			return true
		}
	}
	return false
}

func (p *passportData) passportIDValid() bool {
	return len(p.PID) == 9 &&
		regexp.MustCompile(`[0-9]{9}`).MatchString(p.PID)
}

func validateInRange(num string, start int, end int) bool {
	parsed, err := strconv.Atoi(num)
	return err == nil && start <= parsed && parsed <= end
}

func parseHeight(toParse string) (string, string) {
	height := regexp.MustCompile(`([0-9]*)`).FindString(toParse)
	unit := regexp.MustCompile(`[a-z]{2}`).FindString(toParse)
	return height, unit
}

func setField(v interface{}, name string, value string) error {
	// v must be a pointer to a struct
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return errors.New("v must be pointer to struct")
	}

	// Dereference pointer
	rv = rv.Elem()

	// Lookup field by name
	fv := rv.FieldByName(name)
	if !fv.IsValid() {
		return fmt.Errorf("not a field name: %s", name)
	}

	// Field must be exported
	if !fv.CanSet() {
		return fmt.Errorf("cannot set field %s", name)
	}

	// We expect a string field
	if fv.Kind() != reflect.String {
		return fmt.Errorf("%s is not a string field", name)
	}

	// Set the value
	fv.SetString(value)
	return nil
}

func loadAllPassports(input []string) []passportData {
	var allPassportData []passportData

	currentPassport := passportData{}

	for _, l := range input {

		if strings.TrimSpace(l) == "" {
			allPassportData = append(allPassportData, currentPassport)
			currentPassport = passportData{}
			continue
		}

		fields := strings.Split(l, " ")

		for _, field := range fields {
			keyValue := strings.Split(field, ":")
			setField(&currentPassport, strings.ToUpper(keyValue[0]), keyValue[1])
		}

	}
	allPassportData = append(allPassportData, currentPassport)

	return allPassportData
}

// Part1 of day 3
func (d *Day4Computer) Part1(input shared.Input) (shared.Result, error) {

	validPassports := 0

	for _, passport := range loadAllPassports(input) {

		if passport.hasRequiredFields() {
			validPassports++
		}
	}

	return strconv.Itoa(validPassports), nil
}

// Part2 of day 3
func (d *Day4Computer) Part2(input shared.Input) (shared.Result, error) {
	validPassports := 0

	for _, passport := range loadAllPassports(input) {

		if passport.isValid() {
			validPassports++
		}
	}

	return strconv.Itoa(validPassports), nil
}
