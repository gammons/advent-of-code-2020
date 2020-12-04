package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

func (p *Passport) IsValid() bool {
	return p.validByr() &&
		p.validIyr() &&
		p.validEyr() &&
		p.validHeight() &&
		p.validHairColor() &&
		p.validEyeColor() &&
		p.validPid()
}

func (p *Passport) printInvalid() {
	if !p.validByr() {
		fmt.Printf("Byr is invalid: '%s'\n", p.Byr)
	}
	if !p.validEyr() {
		fmt.Printf("Eyr is invalid: '%s'\n", p.Eyr)
	}
	if !p.validHeight() {
		fmt.Printf("Hgt is invalid: '%s'\n", p.Hgt)
	}
	if !p.validHairColor() {
		fmt.Printf("Hcl is invalid: '%s'\n", p.Hcl)
	}
	if !p.validEyeColor() {
		fmt.Printf("Ecl is invalid: '%s'\n", p.Ecl)
	}
	if !p.validPid() {
		fmt.Printf("Pid is invalid: '%s'\n", p.Pid)
	}
}

func (p *Passport) validByr() bool {
	return p.validYear(p.Byr, 1920, 2002)
}

func (p *Passport) validIyr() bool {
	return p.validYear(p.Iyr, 2010, 2020)
}

func (p *Passport) validEyr() bool {
	return p.validYear(p.Eyr, 2020, 2030)
}

func (p *Passport) validYear(attr string, firstYear int, lastYear int) bool {
	if attr == "" {
		return false
	}
	year, _ := strconv.Atoi(attr)
	return year >= firstYear && year <= lastYear
}

func (p *Passport) validHeight() bool {
	if p.Hgt == "" {
		return false
	}
	r1 := regexp.MustCompile(`\d+`)
	heightNum, _ := strconv.Atoi(r1.FindString(p.Hgt))

	r2 := regexp.MustCompile(`(cm|in)`)
	unit := r2.FindString(p.Hgt)

	if unit == "cm" {
		return heightNum >= 150 && heightNum <= 193
	} else {
		return heightNum >= 59 && heightNum <= 76
	}
}

func (p *Passport) validHairColor() bool {
	if p.Hcl == "" {
		return false
	}
	match, _ := regexp.MatchString(`#(\d|[a-f]){6}`, p.Hcl)
	return match
}

func (p *Passport) validEyeColor() bool {
	if p.Ecl == "" {
		return false
	}
	validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, color := range validColors {
		if p.Ecl == color {
			return true
		}
	}
	return false
}

func (p *Passport) validPid() bool {
	match, _ := regexp.MatchString(`\d{9}`, p.Pid)
	return len(p.Pid) == 9 && match
}

func main() {
	passports := readPassports()
	validCount := 0
	for _, passport := range passports {
		if passport.IsValid() {
			validCount++
		} else {
			// fmt.Println("-----------")
			// passport.printInvalid()
		}
	}
	fmt.Printf("Total valid passports: %d out of %d\n", validCount, len(passports))
}

func readPassports() []*Passport {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var passports []*Passport
	currentPassport := &Passport{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // end of passport
			passports = append(passports, currentPassport)
			currentPassport = &Passport{}
		} else {
			parseLine(line, currentPassport)
		}
	}
	passports = append(passports, currentPassport)

	return passports
}

func parseLine(line string, passport *Passport) {
	attrs := strings.Split(line, " ")
	for _, attr := range attrs {
		splitted := strings.Split(attr, ":")
		key, val := splitted[0], splitted[1]

		switch key {
		case "byr":
			passport.Byr = val
		case "iyr":
			passport.Iyr = val
		case "eyr":
			passport.Eyr = val
		case "hgt":
			passport.Hgt = val
		case "hcl":
			passport.Hcl = val
		case "ecl":
			passport.Ecl = val
		case "pid":
			passport.Pid = val
		case "cid":
			passport.Cid = val
		}
	}
}
