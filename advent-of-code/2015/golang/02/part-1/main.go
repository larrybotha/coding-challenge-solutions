package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getDimStrings(str []byte) [][]byte {
	r, _ := regexp.Compile(`\w+`)
	dims := r.FindAll(str, -1)

	return dims
}

func getDims(str []byte) (int, int, int) {
	r, _ := regexp.Compile(`\d+`)
	dims := r.FindAll(str, -1)
	intDims := make([]int, 3)

	for i, v := range dims {
		intDim, _ := strconv.Atoi(string(v))
		intDims[i] = intDim
	}

	return intDims[0], intDims[1], intDims[2]
}

func getArea(dimsString []byte) int {
	b, w, h := getDims(dimsString)

	fmt.Printf("b: %d, w: %d, h: %d\n", b, w, h)

	return b * w
}

func main() {
	filename := os.Args[1]
	wd, err := os.Getwd()
	paths := []string{wd, filename}

	handleErr(err)

	absFilePath := path.Join(paths...)
	data, err := ioutil.ReadFile(absFilePath)

	handleErr(err)

	dimStrings := getDimStrings(data)

	for _, dimString := range dimStrings {
		getArea(dimString)
	}

	// split text by newline
	// split each line by x
	// create combinations of multiple
	// add half of the smallest
	// sum everything
}
