package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := loadNumbers()
	for idx, num := range numbers {
		for idx2, num2 := range numbers[idx : len(numbers)-1] {
			for _, num3 := range numbers[idx2 : len(numbers)-1] {
				if num+num2+num3 == 2020 {
					fmt.Println("Product = ", num*num2*num3)
					os.Exit(0)
				}
			}
		}
	}
}

func loadNumbers() []int {
	data, _ := ioutil.ReadFile("puzzle_1_input.txt")
	var numbers []int
	for _, n := range strings.Split(string(data), "\n") {
		i, err := strconv.Atoi(n)
		if err == nil {
			numbers = append(numbers, i)
		}
	}
	return numbers
}
