package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// SparkNumbers takes a list of numbers, converts them to floating points and
// generates the sparkline chart from the data. There is a limited number of
// bars provided by the Unicode standard, so there will be the case where
// different numbers are shown in the chart as if they were in the same group.
func SparkNumbers(nums []string) int {
	var output string /* generated sparkline */

	numbers := GetNumbers(nums)
	maxNum := MaxSlice(numbers)
	numLen := len(numbers)
	sparks := make([]rune, 0)
	sticks := []rune{
		'\u2581',
		'\u2582',
		'\u2583',
		'\u2584',
		'\u2585',
		'\u2586',
		'\u2587',
		'\u2588',
	}

	// Get the sparkline for each number.
	for i := 0; i < numLen; i++ {
		unit := int((numbers[i] * 7) / maxNum)
		sparks = append(sparks, sticks[unit])
	}

	// Print the sparklines in the console.
	for i := 0; i < numLen; i++ {
		output += string(sparks[i])
	}

	fmt.Println(output)

	return 0
}

// GetNumbers takes all the string-encoded numbers from the standard input and
// converts them into integers. Non-valid numbers are ignored. Notice that this
// program handles all entries as floating points by default, even if the user
// specifies integers the program will treat them as float64.
func GetNumbers(letters []string) []float64 {
	var numbers []float64

	listlen := len(letters)

	for i := 0; i < listlen; i++ {
		num, err := strconv.ParseFloat(letters[i], 64)

		if err != nil {
			continue
		}

		numbers = append(numbers, num)
	}

	return numbers
}

// MaxSlice returns the maximum number from a list.
func MaxSlice(list []float64) float64 {
	var max float64

	listlen := len(list)

	for i := 0; i < listlen; i++ {
		value := list[i]

		if value > max {
			max = value
		}
	}

	return max
}

func main() {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 1 {
		/* Explode by spaces if they are concatenated */
		lines = strings.Split(lines[0], "\x20")
	}

	if lines == nil {
		fmt.Println("Sparkline")
		fmt.Println("  https://cixtor.com/")
		fmt.Println("  https://github.com/cixtor/sparkline")
		fmt.Println("  https://en.wikipedia.org/wiki/Sparkline")
		fmt.Println("  https://en.wikipedia.org/wiki/Kagi_chart")
		fmt.Println("Usage:")
		fmt.Println("  echo 5 10 22 13 53 | sparkline")
		fmt.Println("  seq 0 50 | sort -R | sparkline")
		fmt.Println("  stock-market +GOOG | sparkline")
		os.Exit(1)
	}

	os.Exit(SparkNumbers(lines))
}
