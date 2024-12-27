package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readLine(f *os.File) (string, error) {
	line := make([]byte, 0)
	buffer := make([]byte, 1)
	var errorToReturn error
	for {
		_, err := f.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				errorToReturn = io.EOF
				break
			} else {
				fmt.Println(err.Error())
				return "", err
			}
		}

		if buffer[0] == byte('\n') {
			break
		} else {
			line = append(line, buffer[0])
		}
	}

	return string(line), errorToReturn
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while opening the file")
		return
	}
	defer f.Close()

	result := make(map[int]bool, 0)
	r := 0

	for i := 0; ; i++ {
		reachedEOF := false
		line, err := readLine(f)
		if err != nil {
			if errors.Is(err, io.EOF) {
				reachedEOF = true
			} else {
				fmt.Println(err.Error())
				return
			}
		}

		tokens := strings.Split(line, ":")
		target, err := strconv.Atoi(tokens[0])
		if err != nil {
			fmt.Printf("Error while converting Atoi: %s", tokens[0])
			return
		}
		operandsString := strings.Split(tokens[1][1:], " ")
		operands := make([]int, 0)
		for _, o := range operandsString {
			operand, err := strconv.Atoi(o)
			if err != nil {
				fmt.Printf("Error while converting Atoi: %s", o)
				return
			}
			operands = append(operands, operand)
		}

		backtrack(i, target, operands, operands[0], 1, result)
		if a, ok := result[i]; ok && a {
			r += target
		}

		if reachedEOF {
			break
		}
	}
	fmt.Println(r)
}

func backtrack(key int, target int, operands []int, s int, index int, result map[int]bool) {
	if index == len(operands) {
		if s == target {
			result[key] = true
		}
		return
	}

	for _, operator := range []rune{'+', '*'} {
		if operator == '+' {
			backtrack(key, target, operands, s+operands[index], index+1, result)
		} else if operator == '*' {
			backtrack(key, target, operands, s*operands[index], index+1, result)
		}
	}
}
