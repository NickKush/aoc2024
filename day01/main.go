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

func parseInputs(filename string) (left_col []int, right_col []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "   ")

		left, _ := strconv.Atoi(data[0])
		left_col = append(left_col, left)

		right, _ := strconv.Atoi(data[1])
		right_col = append(right_col, right)
	}

	return left_col, right_col
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func sort(input []int, low, high int) {
	if low < high {
		p := partition(input, low, high)
		sort(input, low, p-1)
		sort(input, p+1, high)
	}
}

func part1() int {
	left_col, right_col := parseInputs("input.txt")

	sort(left_col, 0, len(left_col)-1)
	sort(right_col, 0, len(right_col)-1)

	result := 0

	for i := 0; i < len(left_col); i++ {
		result += absInt(left_col[i] - right_col[i])
	}

	return result
}

func part2() int {
	left, right := parseInputs("input.txt")

	scores := make(map[int]int)
	for i := 0; i < len(right); i++ {
		scores[right[i]] += 1
	}

	result := 0
	for i := 0; i < len(left); i++ {
		n := left[i]
		result += n * scores[n]
	}

	return result
}

func main() {
	result1 := part1() // Result: 2057374
	fmt.Printf("Result: %d\n", result1)

	result2 := part2() // Result: 23177084
	fmt.Printf("Result: %d\n", result2)
}
