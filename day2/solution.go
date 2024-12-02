package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parseLine(line []byte) ([]int, error) {
	tokens := strings.Split(string(line), " ")
	nums := make([]int, 0)
	for _, tok := range tokens {
		if tok == "" {
			continue
		} else {
			num, err := strconv.Atoi(tok)
			if err != nil {
				return []int{0}, err
			}
			nums = append(nums, num)
		}
	}

	return nums, nil
}

func checkSafety(nums []int) bool {

	var startIsIncreasing bool
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		if i == 1 {
			if diff < 0 {
				startIsIncreasing = false
			} else {
				startIsIncreasing = true
			}
		}

		if startIsIncreasing {
			if diff < 0 {
				return false
			} else if diff > 3 || diff < 1 {
				return false
			}
		} else {
			if diff > 0 {
				return false
			} else if diff < -3 || diff > -1 {
				return false
			}
		}
	}
	return true
}

func main() {
	// Parse file
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while reading the file")
		return
	}

	safeLines := 0

	line := make([]byte, 0)

	buffer := make([]byte, 1)
	var lines [][]int
	lines = make([][]int, 0)
	for {
		_, err := f.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				if len(line) != 0 {
					nums, err := parseLine(line)
					if err != nil {
						fmt.Println(err.Error())
					}

					lines = append(lines, nums)
				}
				break
			} else {
				fmt.Println(err.Error())
				return
			}
		}

		if buffer[0] == byte('\n') {
			nums, err := parseLine(line)
			if err != nil {
				fmt.Println(err.Error())
			}
			lines = append(lines, nums)
			line = make([]byte, 0)
		} else {
			line = append(line, buffer[0])
		}
	}

	for _, nums := range lines {
		isSafe := checkSafety(nums)
		if isSafe {
			safeLines++
		} else {
			for i := range len(nums) {
				sub := make([]int, 0)
				sub = append(sub, nums[:i]...)
				sub = append(sub, nums[i+1:]...)

				if checkSafety(sub) {
					safeLines++
					break
				}
			}
		}
	}

	fmt.Println(safeLines)
}
