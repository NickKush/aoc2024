package main

import (
	_ "embed"
	"fmt"
	"strings"
)

const test_str = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

//go:embed input.txt
var input_str string

func parseInput(value string) (result [][]string) {
	value = strings.TrimRight(value, "\n")
	for _, line := range strings.Split(value, "\n") {
		result = append(result, strings.Split(line, ""))
	}

	return result
}

const expectedString = "XMAS"

// 0 - x, 1 - y
var directions = [...][]int{
	{-1, -1},
	{0, -1},
	{+1, -1},
	{+1, 0},
	{+1, +1},
	{0, +1},
	{-1, +1},
	{-1, 0},
}

func findXmas(arr [][]string, charIndex, lineIndex int, expectedIndex int, dirIndex int) int {
	// fmt.Printf("FindXmas - charIndex: %v, lineIndex: %v, expectedIndex: %v, dirIndex: %v\n", charIndex, lineIndex, expectedIndex, dirIndex)
	result := 0
	for dir_index, dir := range directions {
		if dirIndex >= 0 {
			if dirIndex != dir_index {
				continue
			}
		}

		x := charIndex + dir[0]
		y := lineIndex + dir[1]
		if x < 0 || y < 0 || x > len(arr[lineIndex])-1 || y > len(arr)-1 {
			// fmt.Printf("Boundaries dir_x: %v, dir_y: %v, dir_index: %v\n", dir[0], dir[1], dir_index)
			continue
		}

		char := arr[y][x]

		if char == string(expectedString[expectedIndex]) {
			nextIndex := expectedIndex + 1
			if nextIndex > len(expectedString)-1 {
				result++
				break
			}

			result += findXmas(arr, x, y, nextIndex, dir_index)
		} else {
			continue
		}
	}

	return result
}

func part1() int {
	result := 0
	// data := parseInput(test_str)
	data := parseInput(input_str)

	for line_index, line := range data {
		for char_index, char := range line {
			if char != string(expectedString[0]) {
				continue
			}
			result += findXmas(data, char_index, line_index, 1, -1)
		}
	}

	return result
}

func part2() int {
	result := 0
	// data := parseInput(test_str)
	return result
}

func main() {
	result1 := part1()
	fmt.Printf("Result: %d\n", result1)

	result2 := part2()
	fmt.Printf("Result: %d\n", result2)
}
