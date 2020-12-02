package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Policy struct {
	Char    string
	Minimum int
	Maximum int
}

func (p *Policy) ValidPassword(password string) bool {
	matchCount := 0
	if p.Minimum-1 < len(password) && string(password[p.Minimum-1]) == p.Char {
		matchCount++
	}

	if p.Maximum-1 < len(password) && string(password[p.Maximum-1]) == p.Char {
		matchCount++
	}
	return matchCount == 1
}

func (p *Policy) toString() string {
	return fmt.Sprintf("Char = '%s', Minimum = %d, Maximum = %d", p.Char, p.Minimum, p.Maximum)
}

func main() {
	data := readInput()
	validCount := 0
	for _, input := range data {
		policy := parsePolicy(input)
		if policy.ValidPassword(parsePassword(input)) {
			validCount++
		}
	}
	fmt.Printf("Valid count = %d\n", validCount)
}

func parsePolicy(input string) *Policy {
	r := regexp.MustCompile(`\d+-\d+`)
	minmaxStr := r.FindAllString(input, -1)
	minimum, _ := strconv.Atoi(strings.Split(minmaxStr[0], "-")[0])
	maximum, _ := strconv.Atoi(strings.Split(minmaxStr[0], "-")[1])

	r = regexp.MustCompile(".:")
	char := r.FindString(input)

	return &Policy{
		Char:    string(char[0]),
		Minimum: minimum,
		Maximum: maximum,
	}
}

func parsePassword(input string) string {
	splitted := strings.Split(input, " ")
	return splitted[len(splitted)-1]
}

func readInput() []string {
	data, _ := ioutil.ReadFile("input.txt")
	var input []string

	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			input = append(input, line)
		}
	}
	return input
}
