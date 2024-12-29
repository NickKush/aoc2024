package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

const test_str = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func parseTestStr() (arr [][]int) {
	for _, line := range strings.Split(test_str, "\n") {
		input_data := strings.Split(line, " ")

		levels := []int{}
		for _, v := range input_data {
			n, _ := strconv.Atoi(v)
			levels = append(levels, n)
		}

		arr = append(arr, levels)
	}

	return arr
}

func parseInputs(filename string) (arr [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input_data := strings.Split(scanner.Text(), " ")

		levels := []int{}
		for _, v := range input_data {
			n, _ := strconv.Atoi(v)
			levels = append(levels, n)
		}

		arr = append(arr, levels)
	}

	return arr
}

func isSave(level []int) bool {
	level_type := 0 // -1 - decr , +1 increasing

	for i := 1; i < len(level); i++ {
		a, b := level[i-1], level[i]

		diff := absInt(a - b)
		if diff < 1 || diff > 3 {
			return false
		}

		if a < b { // increasing
			if level_type == 0 {
				level_type = 1
				continue
			}

			if level_type == -1 {
				return false
			}
		} else { // decreasing
			if level_type == 0 {
				level_type = -1
				continue
			}

			if level_type == 1 {
				return false
			}
		}
	}
	return true
}

func part1() int {
	levels := parseInputs("input.txt")
	// levels := parseTestStr()
	// fmt.Printf("Data: %v\n", levels)

	result := 0
	for _, level := range levels {
		if isSave(level) {
			result += 1
		}
	}

	return result
}

func part2() int {
	levels := parseInputs("input.txt")
	// levels := parseTestStr()

	result := 0
	for _, level := range levels {
		safe := isSave(level)
		if safe {
			result += 1
		} else {
			for i := 0; i < len(level); i++ {
				// idk, idc
				t := append([]int(nil), level[:i]...)
				t = append(t, level[i+1:]...)

				if isSave(t) {
					result += 1
					break
				}
			}
		}
	}

	return result
}

func main() {
	result1 := part1()
	fmt.Printf("Result: %d\n", result1)

	result2 := part2()
	fmt.Printf("Result: %d\n", result2)
}
