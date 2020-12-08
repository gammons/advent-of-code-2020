package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Name     string
	Arg      int
	Executed bool
}

type Computer struct {
	Step     int
	Acc      int
	ExitCode int
}

func (c *Computer) Run(instructions []*Instruction) {
	c.Step, c.Acc, c.ExitCode = 0, 0, 0
	for i := 0; i < len(instructions); i++ {
		instructions[i].Executed = false
	}

	for {
		if c.Step == len(instructions) {
			c.ExitCode = 0
			break
		}
		if instructions[c.Step].Executed {
			c.ExitCode = 1
			break
		}
		instructions[c.Step].Executed = true

		switch instructions[c.Step].Name {
		case "acc":
			c.Acc += instructions[c.Step].Arg
			c.Step++
		case "jmp":
			c.Step += instructions[c.Step].Arg
		case "nop":
			c.Step++
		}
	}
}

func main() {
	instructions := readInput()
	comp := &Computer{}

	for i := 0; i < len(instructions); i++ {
		if instructions[i].Name == "jmp" {
			instructions[i].Name = "nop"
			comp.Run(instructions)
			if comp.ExitCode == 0 {
				break
			}
			instructions[i].Name = "jmp"
		}
		if instructions[i].Name == "nop" {
			instructions[i].Name = "jmp"
			comp.Run(instructions)
			if comp.ExitCode == 0 {
				break
			}
			instructions[i].Name = "nop"
		}
	}
	fmt.Printf("acc = %d, exitCode = %d", comp.Acc, comp.ExitCode)
}

func readInput() []*Instruction {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions []*Instruction

	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), " ")
		arg, _ := strconv.Atoi(splitted[1])
		instructions = append(instructions, &Instruction{
			Name: splitted[0],
			Arg:  arg,
		})
	}

	return instructions
}
