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

	result := 0

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

		combinations := make([][]rune, 0)

		backtrack(len(operandsString), []rune{}, &combinations)

		operands := make([]int, 0)
		for _, o := range operandsString {
			operand, err := strconv.Atoi(o)
			if err != nil {
				fmt.Printf("Error while converting Atoi: %s", o)
				return
			}
			operands = append(operands, operand)
		}

		for _, combination := range combinations {
			combinationResult := operands[0]
			for i, operator := range combination {
				if operator == '*' {
					combinationResult *= operands[i+1]
				} else if operator == '+' {
					combinationResult += operands[i+1]
				} else {
					combinationResultString := strconv.Itoa(combinationResult)
					combinationResultString += strconv.Itoa(operands[i+1])
					combinationResult, err = strconv.Atoi(combinationResultString)
					if err != nil {
						fmt.Printf("Error while converting Atoi: %s", combinationResultString)
						return
					}
				}
			}

			if combinationResult == target {
				result += combinationResult
				break
			}
		}

		if reachedEOF {
			break
		}
	}

	fmt.Println(result)
}

func backtrack(operandsLength int, currentCombination []rune, combinations *[][]rune) {
	if len(currentCombination) == operandsLength-1 {
		newSlice := make([]rune, len(currentCombination))
		copy(newSlice, currentCombination)
		*combinations = append(*combinations, newSlice)
		return
	}

	for _, operator := range []rune{'+', '*', '|'} {
		currentCombination = append(currentCombination, operator)
		backtrack(operandsLength, currentCombination, combinations)
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}
