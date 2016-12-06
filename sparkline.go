/**
 * Sparkline
 * https://cixtor.com/
 * https://github.com/cixtor/sparkline
 * https://en.wikipedia.org/wiki/Sparkline
 * https://en.wikipedia.org/wiki/Kagi_chart
 *
 * A sparkline is a very small line chart, typically drawn without axes or
 * coordinates. It presents the general shape of the variation (typically over
 * time) in some measurement, such as temperature or stock market price, in a
 * simple and highly condensed way. Sparklines are small enough to be embedded
 * in text, or several sparklines may be grouped together as elements of a small
 * multiple.
 *
 * Whereas the typical chart is designed to show as much data as possible, and
 * is set off from the flow of text, sparklines are intended to be succinct,
 * memorable, and located where they are discussed.
 *
 * Sparklines are frequently used in line with text. For example: The Dow Jones
 * Industrial Average for February 7, 2006 sparkline which illustrates the
 * fluctuations in the Down Jones index on February 7, 2006. The sparkline
 * should be about the same height as the text around it. Tufte offers some
 * useful design principles for the sizing of sparklines to maximize their
 * readability.
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func DisplayHelp() {
	fmt.Println("Sparkline")
	fmt.Println("  https://cixtor.com/")
	fmt.Println("  https://github.com/cixtor/sparkline")
	fmt.Println("  https://en.wikipedia.org/wiki/Sparkline")
	fmt.Println("  https://en.wikipedia.org/wiki/Kagi_chart")
	fmt.Println("Usage:")
	fmt.Println("  spark [-h|-v] [values]")
	fmt.Println("Examples:")
	fmt.Println("  spark 1 5 22 13 53")
	fmt.Println("  spark $(echo 9 13 5 17 1)")
	fmt.Println("  spark $(seq 0 50 | sort -R)")
}

func SparkNumbers(arguments []string, verbose bool) int {
	var numbers []float64 = GetNumbers(arguments)
	var maxNum float64 = MaxSlice(numbers)
	var numLen int = len(numbers)
	var sparks = make([]rune, 0)

	// Unicode representation of the sparks.
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
		var unit int = int((numbers[i] * 7) / maxNum)
		sparks = append(sparks, sticks[unit])
	}

	if verbose {
		var minNum float64 = MinSlice(numbers)
		var ellipsisLimit int = numLen
		var addThreeDots bool = false
		var ellipsisList []string

		if numLen > 15 {
			ellipsisLimit = 15
			addThreeDots = true
		}

		for j := 0; j < ellipsisLimit; j++ {
			ellipsisList = append(ellipsisList, Float2String(numbers[j]))
		}

		if addThreeDots {
			ellipsisList = append(ellipsisList, "\u2026")
		}

		fmt.Println("Sparkline")
		fmt.Println("Quantity:", numLen)
		fmt.Println("Minimum:", minNum)
		fmt.Println("Maximum:", maxNum)
		fmt.Println("Numbers:", ellipsisList)
	}

	// Print the sparklines in the console.
	var output string
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

func MinSlice(list []float64) float64 {
	var min float64
	var list_len int = len(list)

	for i := 0; i < list_len; i++ {
		var value float64 = list[i]

		if value > min {
			min = value
		}
	}

	return min
}

func Float2String(number float64) string {
	return strconv.FormatFloat(number, 'f', 1, 64)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		DisplayHelp()
		os.Exit(2)
	}

	var arguments []string = os.Args[1:]
	var verbose bool = false

	if os.Args[1] == "-v" {
		arguments = os.Args[2:]
		verbose = true
	}

	os.Exit(SparkNumbers(arguments, verbose))
}
