package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getInputData() []byte {
	wd, err := os.Getwd()
	handleErr(err)

	filepath := os.Args[1:][0]
	paths := []string{wd, filepath}
	absFilePath := path.Join(paths...)

	data, err := ioutil.ReadFile(absFilePath)

	handleErr(err)

	return data
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func minInRange(xs []int) int {
	if len(xs) <= 2 {
		return min(xs[0], xs[1])
	} else {
		return min(xs[0], minInRange(xs[1:]))
	}
}

func getMins(xs []int, n int) []int {
	sides := make([]int, 0)
	// Go's "sort" package allows for sorting of various types
	sort.Ints(xs)

	for i := range xs {
		// append returns a new slice, like [].concat in JS
		sides = append(sides, minInRange(xs[i:]))

		if len(sides) == n {
			break
		}
	}

	return sides
}

func product(xs []int) int {
	total := 1

	for _, x := range xs {
		total *= x
	}

	return total
}

func ribbonLength(xs []int) int {
	shortSides := getMins(xs, 2)
	length := sum(shortSides) * 2
	bowLength := product(xs)

	return length + bowLength
}

func strsToInt(xs []string) []int {
	ints := make([]int, 0)

	for _, x := range xs {
		// to convert from alpha to int, use strconv.Atoi
		n, _ := strconv.Atoi(x)
		ints = append(ints, n)
	}

	return ints
}

func sum(xs []int) int {
	total := 0

	for _, x := range xs {
		total += x
	}

	return total
}

func main() {
	data := getInputData()
	// regular expressions don't use the /[expr]/ format in JS and PHP
	//
	// regular expression patterns can be written using backticks, or double quotes.
	// double-quotes are interpreted string literals, and characters such as backslashes
	// need to be escaped

	// rx := regexp.MustCompile("\\w+")
	rx := regexp.MustCompile(`\w+`)

	// Regex.FindAll requires an int value to determine where to stop searching
	// -1 indicates search the entire string
	lines := rx.FindAll(data, -1)
	dims := make([][]int, 0)

	for _, v := range lines {
		// the "strings" package allows for string manipulation, such as splitting a string
		giftDims := strings.Split(string(v), `x`)
		intGiftDims := strsToInt(giftDims)
		dims = append(dims, intGiftDims)
	}

	lengths := make([]int, 0)

	for _, x := range dims {
		lengths = append(lengths, ribbonLength(x))
	}

	totalLength := sum(lengths)

	fmt.Printf("total ribbon length: %d\n", totalLength)
}
