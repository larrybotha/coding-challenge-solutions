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
	inputFilename := os.Args[1:][0]
	fmt.Print(inputFilename)

	dirname, err := os.Getwd()
	check(err)

	fmt.Print(dirname)

	absFilePath := filepath.Join(dirname, inputFilename)

	dat, err := ioutil.ReadFile(absFilePath)
	check(err)

	fmt.Print(dat)
}
