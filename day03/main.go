package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"unicode"
)

const (
	test_str   = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	test_str_2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
)

//go:embed input.txt
var input_str string

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic("oh no")
	}

	return result
}

func part1() int {
	result := 0

	// data := []rune(test_str)
	data := []rune(input_str)

	str_len := len(data)
	i := 0

	for i <= str_len {
		if (str_len - i) < 3 {
			break
		}

		is_mul := string(data[i : i+3])
		if is_mul == "mul" {
			i += 2

			token := data[i+1]
			if token != '(' {
				goto cont
			}
			i++

			// search for numbers
			var first_digit, second_digit string
			step := 0
			for {
				token = data[i+1]
				if unicode.IsDigit(token) {
					if step == 0 {
						first_digit += string(token)
					} else {
						second_digit += string(token)
					}
					i++
					continue
				}

				if token == ',' {
					i++
					step = 1
					continue
				}

				if token == ')' {
					i++
					break
				}

				// error
				goto cont
			}

			result += toInt(first_digit) * toInt(second_digit)
		}
	cont:
		i++
	}

	return result
}

func part2() int {
	result := 0

	// data := []rune(test_str_2)
	data := []rune(input_str)

	str_len := len(data)
	i := 0

	is_enabled := true

	for i <= str_len {
		if (str_len - i) < 3 {
			break
		}

		is_do := string(data[i : i+4])
		if is_do == "do()" {
			is_enabled = true
		}
		is_dont := string(data[i : i+7])
		if is_dont == "don't()" {
			is_enabled = false
		}

		is_mul := string(data[i : i+3])
		if is_mul == "mul" {
			if !is_enabled {
				goto cont
			}

			i += 2

			token := data[i+1]
			if token != '(' {
				goto cont
			}
			i++

			// search for numbers
			var first_digit, second_digit string
			step := 0
			for {
				token = data[i+1]
				if unicode.IsDigit(token) {
					if step == 0 {
						first_digit += string(token)
					} else {
						second_digit += string(token)
					}
					i++
					continue
				}

				if token == ',' {
					i++
					step = 1
					continue
				}

				if token == ')' {
					i++
					break
				}

				// error
				goto cont
			}

			result += toInt(first_digit) * toInt(second_digit)
		}
	cont:
		i++
	}

	return result
}

func main() {
	result1 := part1()
	fmt.Printf("Result: %d\n", result1)

	result2 := part2()
	fmt.Printf("Result: %d\n", result2)
}
