package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type Point struct {
	guard rune
	x     int
	y     int
}

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

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while opening the file")
		return
	}
	defer f.Close()

	m := make([][]rune, 0)

	var guard rune
	guardCoordinates := make([]int, 2)

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

		row := make([]rune, 0)

		for j, b := range line {
			if b == '^' || b == '>' || b == '<' || b == 'v' {
				guardCoordinates = []int{i, j}
				guard = rune(b)
			}
			row = append(row, rune(b))
		}
		m = append(m, row)

		if reachedEOF {
			break
		}
	}

	rows := len(m)
	cols := len(m[0])

	i, j := guardCoordinates[0], guardCoordinates[1]
	visited := make(map[Point]bool, 0)

	result := 0

	for i < rows && i >= 0 && j < cols && j >= 0 {
		point := Point{'-', i, j}
		if _, ok := visited[point]; !ok {
			result++
			visited[point] = true
		}

		if guard == '^' {
			if i == 0 {
				break
			}

			if m[i-1][j] == '#' {
				guard = '>'
			} else {
				i--
			}
		} else if guard == '>' {
			if j == cols-1 {
				break
			}

			if m[i][j+1] == '#' {
				guard = 'v'
			} else {
				j++
			}
		} else if guard == '<' {
			if j == 0 {
				break
			}

			if m[i][j-1] == '#' {
				guard = '^'
			} else {
				j--
			}
		} else if guard == 'v' {
			if i == rows-1 {
				break
			}

			if m[i+1][j] == '#' {
				guard = '<'
			} else {
				i++
			}
		}
	}

	fmt.Println(result)

	obstacles := 0

	for point, _ := range visited {
		prev := m[point.x][point.y]
		m[point.x][point.y] = '#'
		visitedInside := make(map[Point]bool, 0)

		i, j := guardCoordinates[0], guardCoordinates[1]
		guard = '^'

		for i < rows && i >= 0 && j < cols && j >= 0 {
			if _, ok := visitedInside[Point{guard, i, j}]; ok {
				obstacles++
				break
			}
			visitedInside[Point{guard, i, j}] = true

			if guard == '^' {
				if i == 0 {
					break
				}

				if m[i-1][j] == '#' {
					guard = '>'
				} else {
					i--
				}
			} else if guard == '>' {
				if j == cols-1 {
					break
				}

				if m[i][j+1] == '#' {
					guard = 'v'
				} else {
					j++
				}
			} else if guard == '<' {
				if j == 0 {
					break
				}

				if m[i][j-1] == '#' {
					guard = '^'
				} else {
					j--
				}
			} else if guard == 'v' {
				if i == rows-1 {
					break
				}

				if m[i+1][j] == '#' {
					guard = '<'
				} else {
					i++
				}
			}
		}
		m[point.x][point.y] = prev
	}

	fmt.Println(obstacles)

}
