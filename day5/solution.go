package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readLine(f *os.File) ([]byte, error) {
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
				return nil, err
			}
		}

		if buffer[0] == byte('\n') {
			break
		} else {
			line = append(line, buffer[0])
		}
	}

	return line, errorToReturn
}

func parseRule(rule []byte) (string, string) {
	tokens := strings.Split(string(rule), "|")
	return tokens[0], tokens[1]
}

func parseUpdate(update []byte) []string {
	return strings.Split(string(update), ",")
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while opening the file")
		return
	}
	defer f.Close()

	rules := make(map[string][]string, 0)

	// Parse rules
	for {
		ruleLine, err := readLine(f)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if len(ruleLine) == 0 {
			break
		}
		start, end := parseRule(ruleLine)
		if _, ok := rules[start]; !ok {
			rules[start] = make([]string, 0)
		}

		rules[start] = append(rules[start], end)
	}

	correctUpdates := make([][]string, 0)
	incorrectUpdates := make([][]string, 0)
	// Parse lines
	for {
		reachedEOF := false
		updateLine, err := readLine(f)
		if err != nil {
			if errors.Is(err, io.EOF) {
				reachedEOF = true
			} else {
				fmt.Println(err.Error())
				return
			}
		}

		updates := parseUpdate(updateLine)

		isCorrect := true
		visited := make(map[string]bool, 0)
		for _, update := range updates {
			visited[update] = true

			for _, end := range rules[update] {
				if _, ok := visited[end]; ok {
					isCorrect = false
				}
			}
		}

		if isCorrect {
			correctUpdates = append(correctUpdates, updates)
		} else {
			incorrectUpdates = append(incorrectUpdates, updates)
		}

		if reachedEOF {
			break
		}
	}

	sum := 0
	for _, update := range correctUpdates {
		middleUpdate := update[len(update)/2]
		middle, err := strconv.Atoi(middleUpdate)
		if err != nil {
			fmt.Println("Error while converting to int")
			return
		}
		sum += middle
	}
	fmt.Println(sum)

	incorrectSum := 0
	for _, update := range incorrectUpdates {
		h := make(map[string]bool, 0)

		for _, u := range update {
			h[u] = true
		}

		// Calculate the topological order
		order := make([]string, 0)
		queue := make([]string, 0)
		inDegree := make(map[string]int, 0)
		for _, end := range rules {
			for _, node := range end {
				if _, ok := inDegree[node]; !ok {
					inDegree[node] = 0
				}
				if _, ok := h[node]; ok {
					inDegree[node]++
				}
			}
		}

		for i := range 100 {
			a := strconv.Itoa(i)
			if _, ok := inDegree[a]; !ok || inDegree[a] == 0 {
				queue = append(queue, a)
			}
		}

		for len(queue) != 0 {
			order = append(order, queue[0])
			end := rules[queue[0]]
			for _, node := range end {
				inDegree[node]--
				if in, ok := inDegree[node]; ok && in == 0 {
					queue = append(queue, node)
				}
			}
			queue = queue[1:]
		}

		orderedUpdate := make([]string, 0)

		for _, orderElement := range order {
			if _, ok := h[orderElement]; ok {
				orderedUpdate = append(orderedUpdate, orderElement)
			}
		}

		middle, _ := strconv.Atoi(orderedUpdate[len(orderedUpdate)/2])
		incorrectSum += middle
	}

	fmt.Println(incorrectSum)
}
