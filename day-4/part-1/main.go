package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	return p.Byr != "" &&
		p.Iyr != "" &&
		p.Eyr != "" &&
		p.Hgt != "" &&
		p.Hcl != "" &&
		p.Ecl != "" &&
		p.Pid != ""
}

func main() {
	passports := readPassports()
	validCount := 0
	for _, passport := range passports {
		if passport.IsValid() {
			validCount++
		}
	}
	fmt.Println("Total valid passports: ", validCount)
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
