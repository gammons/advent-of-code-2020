package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	Answers map[string]bool
}

type Group struct {
	People []*Person
}

func (g *Group) AnswerCount() int {
	questions := "abcdefghijklmnopqrstuvwxyz"
	groupAnswers := make(map[string]bool)

	for _, question := range strings.Split(questions, "") {
		groupAnswer := true
		for _, person := range g.People {
			groupAnswer = person.Answers[question] && groupAnswer
		}
		if groupAnswer {
			groupAnswers[question] = groupAnswer
		}
	}
	return len(groupAnswers)
}

func NewPerson() *Person {
	return &Person{
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
	currentGroup := &Group{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			groups = append(groups, currentGroup)
			currentGroup = &Group{}
		} else {
			person := NewPerson()
			currentGroup.People = append(currentGroup.People, person)
			parseAnswer(line, person)
		}
	}
	groups = append(groups, currentGroup)

	return groups
}

func parseAnswer(line string, person *Person) {
	for _, char := range strings.Split(line, "") {
		person.Answers[char] = true
	}
}
