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

func shortestSides(xs []int) []int {
	sides := make([]int, 0)
	sort.Ints(xs)

	for i := range xs {
		sides = append(sides, minInRange(xs[i:]))

		if len(sides) == 2 {
			break
		}
	}

	return sides
}

func sliceProduct(xs []int) int {
	product := 1

	for _, x := range xs {
		product *= x
	}

	return product
}

func ribbonLength(xs []int) int {
	shortSides := shortestSides(xs)
	length := sum(shortSides) * 2
	bowLength := sliceProduct(xs)

	return length + bowLength
}

func split(str, match string) []string {
	return strings.Split(str, match)
}

func strsToInt(xs []string) []int {
	ints := make([]int, 0)

	for _, x := range xs {
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
	/**
	 * regular expressions don't use the /[expr]/ format in JS and PHP
	 */

	rx := regexp.MustCompile(`\w+`)
	lines := rx.FindAll(data, -1)
	dims := make([][]int, 0)

	for _, v := range lines {
		giftDims := split(string(v), `x`)
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
