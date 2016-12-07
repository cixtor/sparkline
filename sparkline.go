package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SparkNumbers(nums []string) int {
	var output string /* generated sparkline */
	var numbers []float64 = GetNumbers(nums)
	var maxNum float64 = MaxSlice(numbers)
	var numLen int = len(numbers)
	var sparks = make([]rune, 0)
	var sticks = []rune{
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

func GetNumbers(letters []string) []float64 {
	var numbers []float64
	var list_len int = len(letters)

	for i := 0; i < list_len; i++ {
		num, _ := strconv.ParseFloat(letters[i], 64)
		numbers = append(numbers, num)
	}

	return numbers
}

func MaxSlice(list []float64) float64 {
	var max float64
	var list_len int = len(list)

	for i := 0; i < list_len; i++ {
		var value float64 = list[i]

		if value > max {
			max = value
		}
	}

	return max
}

func main() {
	var lines []string
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

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
