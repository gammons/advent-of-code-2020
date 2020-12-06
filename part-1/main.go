package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Group struct {
	Answers map[string]bool
}

func (g *Group) AnswerCount() int {
	return len(g.Answers)
}

func NewGroup() *Group {
	return &Group{
		Answers: make(map[string]bool),
	}
}

func main() {
	groups := readGroups()
	countSum := 0
	for _, group := range groups {
		countSum += group.AnswerCount()
	}
	fmt.Println("Sum: ", countSum)
}

func readGroups() []*Group {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var groups []*Group
	currentGroup := NewGroup()

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			groups = append(groups, currentGroup)
			currentGroup = NewGroup()
		} else {
			parseLine(line, currentGroup)
		}
	}
	groups = append(groups, currentGroup)

	return groups
}

func parseLine(line string, group *Group) {
	for _, char := range strings.Split(line, "") {
		group.Answers[char] = true
	}
}
