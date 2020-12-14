package main

import (
	"bufio"
	"regexp"

	"strings"

	"fmt"
	"log"
	"os"
	"strconv"
)

type BitNum string

func CreateBitNumFromInt(n int) BitNum {
	stringBitNum := BitNum(strconv.FormatUint(uint64(n), 2)).reverse()
	for i := len(stringBitNum); i < 36; i++ {
		stringBitNum += "0"
	}
	return stringBitNum.reverse()
}

func (b BitNum) ApplyMask(mask BitNum) BitNum {
	out := []rune(b)
	for i := range b {
		if string(mask[i]) == "X" {
			continue
		}
		out[i] = rune(mask[i])
	}
	return BitNum(out)
}

func (b BitNum) ToInt() int {
	i, _ := strconv.ParseInt(string(b), 2, 64)
	return int(i)
}

func (m BitNum) reverse() BitNum {
	runes := []rune(m)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return BitNum(runes)
}

func main() {
	var currentMask BitNum
	memory := make(map[int]int)

	r := regexp.MustCompile(`\d+`)
	for _, line := range readInput() {
		if strings.HasPrefix(line, "mask") {
			currentMask = BitNum(strings.Split(line, " = ")[1])
		} else {
			matches := r.FindAllString(line, -1)
			addr, _ := strconv.Atoi(matches[0])
			val, _ := strconv.Atoi(matches[1])
			bVal := CreateBitNumFromInt(val)
			masked := bVal.ApplyMask(currentMask)

			memory[addr] = masked.ToInt()
		}
	}

	sum := 0
	for _, mem := range memory {
		sum += mem
	}
	fmt.Println("Sum = ", sum)
}

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output []string
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	return output
}
