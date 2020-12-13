package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Schedule struct {
	Timestamp int
	Buses     []int
}

func (s *Schedule) WaitMinutes() int {
	earliest := 9999
	var earliestBus int

	for _, bus := range s.Buses {
		next := bus - (s.Timestamp % bus)
		if next < earliest {
			earliest = next
			earliestBus = bus
		}
	}
	return earliestBus * earliest
}

func main() {
	schedule := readInput()
	fmt.Println("Earlist: ", schedule.WaitMinutes())
}

func readInput() *Schedule {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timestamp, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	busInput := strings.Split(scanner.Text(), ",")

	var buses []int

	for _, bus := range busInput {
		if bus == "x" {
			continue
		}
		busID, _ := strconv.Atoi(bus)
		buses = append(buses, busID)
	}
	return &Schedule{Timestamp: timestamp, Buses: buses}
}
