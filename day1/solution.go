package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseLine(line []byte) (int, int, error) {
	tokens := strings.Split(string(line), " ")
	nums := make([]int, 2)
	i := 0
	for _, tok := range tokens {
		if tok == "" {
			continue
		} else {
			num, err := strconv.Atoi(tok)
			if err != nil {
				return 0, 0, err
			}
			nums[i] = num
			i++
		}
	}

	return nums[0], nums[1], nil
}

func handleLine(locations1 *[]int, locations2 *[]int, line []byte) error {
	num1, num2, err := parseLine(line)
	if err != nil {
		return err
	}
	*locations1 = append(*locations1, num1)
	*locations2 = append(*locations2, num2)
	return nil
}

func main() {
	// Parse input
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while reading the file")
		return
	}
	defer f.Close()

	locations1 := make([]int, 0)
	locations2 := make([]int, 0)

	line := make([]byte, 0)

	buffer := make([]byte, 1)
	for {
		_, err := f.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				if len(line) != 0 {
					err = handleLine(&locations1, &locations2, line)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
				}
				break
			} else {
				fmt.Println(err.Error())
				return
			}
		}

		if buffer[0] == byte('\n') {
			err = handleLine(&locations1, &locations2, line)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			line = make([]byte, 0)
		} else {
			line = append(line, buffer[0])
		}
	}

	// Part 1

	slices.Sort(locations1)
	slices.Sort(locations2)

	diff := 0

	for i := range len(locations1) {
		diff += int(math.Abs(float64(locations1[i] - locations2[i])))
	}
	fmt.Println(diff)

	// Part two
	h := make(map[int]int, 0)

	similarity := 0

	for _, num := range locations2 {
		if _, ok := h[num]; !ok {
			h[num] = 0
		}
		h[num]++
	}

	for _, num := range locations1 {
		if _, ok := h[num]; ok {
			similarity += num * h[num]
		}
	}

	fmt.Println(similarity)

}
