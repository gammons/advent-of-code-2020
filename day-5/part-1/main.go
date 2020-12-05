package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const Rows = 127
const Columns = 7

type BoardingPass struct {
	input string
}

func (b *BoardingPass) Row() int {
	low := 0
	high := Rows
	for _, char := range b.input[:7] {
		mid := ((high - low) / 2) + 1
		if string(char) == "B" {
			low = low + mid
		} else {
			high = high - mid
		}
	}
	return low
}

func (b *BoardingPass) Column() int {
	low := 0
	high := Columns
	for _, char := range b.input[7:] {
		mid := ((high - low) / 2) + 1
		if string(char) == "R" {
			low = low + mid
		} else {
			high = high - mid
		}
		//fmt.Printf("char = %s, low = %d, high = %d, mid = %d\n", string(char), low, high, mid)
	}
	return low
}

func (b *BoardingPass) SeatID() int {
	return b.Row()*8 + b.Column()
}

func main() {
	maxID := 0
	for _, pass := range readInput() {
		id := pass.SeatID()
		if id > maxID {
			maxID = id
		}
	}

	fmt.Println("Max ID: ", maxID)
}

func readInput() []*BoardingPass {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var boardingPasses []*BoardingPass

	for scanner.Scan() {
		line := scanner.Text()
		boardingPasses = append(boardingPasses, &BoardingPass{input: line})
	}

	return boardingPasses
}
