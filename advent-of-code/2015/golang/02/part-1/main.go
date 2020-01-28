package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	filename := os.Args[1]
	wd, err := os.Getwd()
	paths := []string{wd, filename}

	handleErr(err)

	absFilePath := path.Join(paths...)

	data, err := ioutil.ReadFile(absFilePath)

	handleErr(err)

	fmt.Printf("filename: %s\n", data)

	// split text by newline
	// split each line by x
	// create combinations of multiple
	// add half of the smallest
	// sum everything
}
