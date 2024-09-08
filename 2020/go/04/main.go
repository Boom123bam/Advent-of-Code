package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Passport [8]string

func countFields(p Passport) int {
	count := 0
	for _, fieldVal := range p {
		if fieldVal != "" {
			count++
		}
	}
	return count
}

func isVaild(p Passport) bool {
	byr, err1 := strconv.Atoi(p[0])
	iyr, err2 := strconv.Atoi(p[1])
	eyr, err3 := strconv.Atoi(p[2])
	hgt, err4 := strconv.Atoi(p[3][:len(p[3])-2])
	hgtUnit := p[3][len(p[3])-2:]
	hcl, ecl, pid := p[4], p[5], p[6]
	// check units
	if !(hgtUnit == "cm" || hgtUnit == "in") {
		return false
	}
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return false
	}
	// byr at least 1920 and at most 2002
	if (byr < 1920 || byr > 2002) ||
		// iyr at least 2010 and at most 2020
		(iyr < 2010 || iyr > 2020) ||
		// eyr at least 2020 and at most 2030
		(eyr < 2020 || eyr > 2030) ||
		// If cm, the number must be at least 150 and at most 193.
		(hgtUnit == "cm" && (hgt < 150 || hgt > 193)) ||
		// If in, the number must be at least 59 and at most 76
		(hgtUnit == "in" && (hgt < 59 || hgt > 76)) {
		return false
	}
	// hcl is a # followed by exactly six characters 0-9 or a-f
	if hcl[0] != '#' || len(hcl) != 7 {
		return false
	}
	for i := 1; i < 6; i++ {
		c := hcl[i]
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return false
		}
	}
	// ecl is exactly one of amb blu brn gry grn hzl oth
	if !(ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth") ||
		// pid is a nine-digit number, including leading zeroes
		len(pid) != 9 {
		return false
	}
	return true
}

func main() {
	file, err := os.Open("04/input.txt")
	// file, err := os.Open("04/test.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	passportFields := [8]string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		"cid",
	}
	fieldToIndex := make(map[string]int)

	for i, field := range passportFields {
		fieldToIndex[field] = i
	}

	var passports []Passport

	passportIndex := 0
	passports = append(passports, Passport{})

	// scan lines
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			// empty line
			passportIndex++
			passports = append(passports, Passport{})
			continue
		}

		for _, field := range fields {
			data := strings.Split(field, ":")
			fieldName := data[0]
			fieldValue := data[1]
			passports[passportIndex][fieldToIndex[fieldName]] = fieldValue
		}
	}
	// remove last empty item
	passports = passports[:len(passports)-1]

	// part 1
	p1Count := 0
	p2Count := 0
	for _, passport := range passports {
		switch countFields(passport) {
		case 8:
			break
		case 7:
			if passport[fieldToIndex["cid"]] == "" {
				break
			}
			continue
		default:
			continue
		}
		p1Count++
		if isVaild(passport) {
			p2Count++
		}

	}
	fmt.Println(p1Count)
	fmt.Println(p2Count)
}
