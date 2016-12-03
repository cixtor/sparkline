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

import "os"
import "fmt"
import "strconv"

func main() {
	if len(os.Args) > 1 {
		var numbers []float64 = get_numbers(os.Args)
		var max_num float64 = max_slice(numbers)
		// var min_num int = min_slice(numbers)
		var num_len int = len(numbers)
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
		for i := 0; i < num_len; i++ {
			var unit int = int((numbers[i] * 7) / max_num)
			sparks = append(sparks, sticks[unit])
		}

		// Print the sparklines in the console.
		for i := 0; i < num_len; i++ {
			fmt.Printf("%s", string(sparks[i]))
		}

		fmt.Printf("\n")

	} else {
		os.Exit(1)
	}
}

func get_numbers(arguments []string) []float64 {
	var numbers []float64
	var letters []string = arguments[1:]
	var list_len int = len(letters)

	for i := 0; i < list_len; i++ {
		num, _ := strconv.ParseFloat(letters[i], 64)
		numbers = append(numbers, num)
	}

	return numbers
}

func max_slice(list []float64) float64 {
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

func min_slice(list []float64) float64 {
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
