package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	Name  string
	Holds []*BagCount
}

type BagCount struct {
	Bag   *Bag
	Count int
}

func (b *Bag) InnerCount() int {
	count := 0
	for _, bag := range b.Holds {
		count += bag.Count
		innerCount := bag.Bag.InnerCount()
		if innerCount > 0 {
			count += (bag.Count * innerCount)
		}
	}
	return count
}

func main() {
	bagHash := readInput()
	fmt.Println("count = ", bagHash["shiny gold"].InnerCount())
}

func readInput() map[string]*Bag {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bags := make(map[string]*Bag)

	for scanner.Scan() {
		line := scanner.Text()
		parse(bags, line)
	}

	return bags
}

func parse(bags map[string]*Bag, line string) {
	name := strings.Join(strings.Split(line, " ")[:2], " ")
	if bags[name] == nil {
		bags[name] = &Bag{Name: name}
	}

	if strings.Contains(line, "no other bags") {
		return
	}

	contains := strings.Split(strings.Split(line, "contain")[1], " ")

	for i := 0; i < len(contains)-1; i += 4 {
		count, _ := strconv.Atoi(contains[i+1])
		containsName := fmt.Sprintf("%s %s", contains[i+2], contains[i+3])

		if bags[containsName] == nil {
			bags[containsName] = &Bag{Name: containsName}
		}

		bags[name].Holds = append(bags[name].Holds, &BagCount{
			Count: count,
			Bag:   bags[containsName],
		})
	}
}
