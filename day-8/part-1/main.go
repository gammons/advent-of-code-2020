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

func main() {
	instructions := readInput()
	step, acc := 0, 0
	for {
		if instructions[step].Executed {
			break
		}
		instructions[step].Executed = true

		switch instructions[step].Name {
		case "acc":
			acc += instructions[step].Arg
			step++
		case "jmp":
			step += instructions[step].Arg
		case "nop":
			step++
		}
	}
	fmt.Println("acc = ", acc)
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
