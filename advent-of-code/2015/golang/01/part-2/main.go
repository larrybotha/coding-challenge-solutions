package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

/**
 * From the golang docs:
 * The error built-in interface type is the conventional interface for representing
 * an error condition, with the nil value representing no error.
 */
func handleErr(err error) {
	// check for errors using nil
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	inputFilename := os.Args[1:][0]
	dir, err := os.Getwd()
	handleErr(err)

	absFilePath := path.Join(dir, inputFilename)

	// open file
	data, err := ioutil.ReadFile(absFilePath)

	/**
	 * instead of using fmt.Println(string(data)), fmt.Printf will interpret data as a string
	 */
	// fmt.Println(data)
	fmt.Printf("data: %s", data)

	handleErr(err)

	floor := 0
	pos := 0

	for i, v := range data {
		if v == ')' {
			floor--
		} else if v == '(' {
			floor++
		}

		if floor == -1 {
			pos = i + 1
			break
		}
	}

	fmt.Printf("pos: %d\n", pos)
}
