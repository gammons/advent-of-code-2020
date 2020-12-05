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
	}
	return low
}

func (b *BoardingPass) SeatID() int {
	return b.Row()*8 + b.Column()
}

func main() {
	seats := make([][]int, Rows+1)
	for i := range seats {
		seats[i] = make([]int, Columns+1)
	}

	for _, pass := range readInput() {
		seats[pass.Row()][pass.Column()] = pass.SeatID()
	}

	type EmptySeat struct {
		Row    int
		Column int
	}

	var emptySeats []*EmptySeat

	for row := range seats {
		if row == 0 || row == 127 {
			continue
		}
		for col := range seats[row] {
			if seats[row][col] == 0 {
				emptySeats = append(emptySeats, &EmptySeat{Row: row, Column: col})
			}
		}
	}

	fmt.Println("Empty seats:")
	for _, seat := range emptySeats {
		id := seat.Row*8 + seat.Column
		fmt.Printf("Row: %d, Col: %d, ID: %d\n", seat.Row, seat.Column, id)
	}
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
