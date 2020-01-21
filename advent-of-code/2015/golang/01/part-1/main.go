package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	/**
	 * get command line args... from index 1?
	 */
	inputFilename := os.Args[1:][0]
	dirname, err := os.Getwd()
	check(err)
	absFilePath := filepath.Join(dirname, inputFilename)

	/**
	 * get a []byte of the file contents
	 */
	dat, err := ioutil.ReadFile(absFilePath)
	check(err)

	/**
	 * we can parse the []byte to a string
	 */
	fmt.Print(string(dat))

	count := 0

	for _, c := range string(dat) {
		if c == '(' {
			count++
		}

		if c == ')' {
			count--
		}
	}

	fmt.Printf("Floor: %d\n", count)
}
