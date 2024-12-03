package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	buffer := make([]byte, 1)

	doing := true
	parsingDo := false
	isDoPath := true
	latest_char_do := 'a'

	parsing := false
	total_sum := 0
	latest_char := 'a'
	parsing_first_number := false
	parsing_second_number := false

	var first_num_s string
	var second_num_s string

	var first_num int
	var second_num int

	for {
		_, err := f.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				fmt.Println(err.Error())
				return
			}
		}

		c := rune(buffer[0])

		if !parsingDo {
			if c == 'd' {
				parsingDo = true
				latest_char_do = 'd'
			}
		} else {
			if latest_char_do == 'd' && c == 'o' {
				latest_char_do = 'o'
			} else if latest_char_do == 'o' && c == '(' {
				latest_char_do = '('
				isDoPath = true
			} else if latest_char_do == '(' && c == ')' {
				doing = isDoPath
				parsingDo = false
				latest_char_do = 'a'
			} else if latest_char_do == 'o' && c == 'n' {
				latest_char_do = 'n'
				isDoPath = false
			} else if latest_char_do == 'n' && c == '\'' {
				latest_char_do = '\''
			} else if latest_char_do == '\'' && c == 't' {
				latest_char_do = 't'
			} else if latest_char_do == 't' && c == '(' {
				latest_char_do = '('
			} else {
				parsingDo = false
				latest_char_do = 'a'
			}
		}

		if doing {
			if !parsing {
				if c == 'm' {
					parsing = true
					latest_char = 'm'
				}
			} else {
				if latest_char == 'm' && c == 'u' {
					latest_char = 'u'
				} else if latest_char == 'u' && c == 'l' {
					latest_char = 'l'
				} else if latest_char == 'l' && c == '(' {
					latest_char = '('
				} else if latest_char == '(' {
					if !parsing_first_number && !parsing_second_number {
						if unicode.IsDigit(c) {
							parsing_first_number = true
							first_num_s += string(c)
						} else {
							parsing_first_number = false
							parsing_second_number = false
							parsing = false
							first_num_s = ""
							second_num_s = ""
							latest_char = 'a'
						}
					} else if parsing_first_number && !parsing_second_number {
						if unicode.IsDigit(c) {
							first_num_s += string(c)
						} else {
							if c == ',' {
								num, err := strconv.Atoi(first_num_s)
								if err != nil {
									fmt.Println("Error while converting first num: " + err.Error())
									return
								}
								first_num = num

								parsing_first_number = false
								parsing_second_number = true
							} else {
								parsing_first_number = false
								parsing_second_number = false
								parsing = false
								first_num_s = ""
								second_num_s = ""
								latest_char = 'a'
							}
						}
					} else if !parsing_first_number && parsing_second_number {
						if unicode.IsDigit(c) {
							second_num_s += string(c)
						} else {
							if c == ')' {
								num, err := strconv.Atoi(second_num_s)
								if err != nil {
									fmt.Println("Error while converting second num: " + err.Error())
									return
								}
								second_num = num
								total_sum += first_num * second_num
								parsing_first_number = false
								parsing_second_number = false
								parsing = false
								first_num_s = ""
								second_num_s = ""
								latest_char = 'a'

							} else {
								parsing_first_number = false
								parsing_second_number = false
								parsing = false
								first_num_s = ""
								second_num_s = ""
								latest_char = 'a'
							}
						}

					}
				}
			}
		}
	}
	fmt.Println(total_sum)
}
