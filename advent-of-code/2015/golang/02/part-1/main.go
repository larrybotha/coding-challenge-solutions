package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func splitByNewLines(str []byte) [][]byte {
	rx := regexp.MustCompile(`\w+`)
	lines := rx.FindAll(str, -1)

	return lines
}

func getDims(bs []byte) []int {
	/**
	* We could use a regex, but strings.Slice is a nice and simple way to do this
	* see https://yourbasic.org/golang/string-functions-reference-cheat-sheet/
	 */
	// r, _ := regexp.Compile(`\d+`)
	// dims := r.FindAll(str, -1)
	dims := strings.Split(string(bs), "x")
	intDims := make([]int, 3)

	for i, v := range dims {
		intDim, _ := strconv.Atoi(string(v))
		intDims[i] = intDim
	}

	return intDims
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func minInSlice(xs []int) int {
	if len(xs) < 2 {
		panic("Too few values in slice")
	}

	if len(xs) == 2 {
		return min(xs[0], xs[1])
	}

	return min(xs[0], minInSlice(xs[1:]))
}

func sum(xs []int) int {
	total := 0

	for _, v := range xs {
		total += v
	}

	return total
}

func getCartesianProductSlice(xs []int) []int {
	ys := make([]int, 0)

	for i, v := range xs {
		for j, v2 := range xs {
			if i != j {
				ys = append(ys, v*v2)
			}
		}
	}

	return ys
}

func getTotalArea(dimsString []byte) int {
	dims := getDims(dimsString)
	areas := getCartesianProductSlice(dims)
	total := sum(areas) + minInSlice(areas)

	return total
}

func main() {
	filename := os.Args[1]
	wd, err := os.Getwd()
	paths := []string{wd, filename}

	handleErr(err)

	absFilePath := path.Join(paths...)
	data, err := ioutil.ReadFile(absFilePath)

	handleErr(err)

	dimStrings := splitByNewLines(data)
	totalArea := 0

	for _, dimString := range dimStrings {
		totalArea += getTotalArea(dimString)
	}

	fmt.Printf("total area: %d\n", totalArea)
}
