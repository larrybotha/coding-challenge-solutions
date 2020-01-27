package main

import (
	"fmt"
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
		panic(err)
	}
}

func main() {
	inputFilename := os.Args[1:][0]
	dir, err := os.Getwd()
	handleErr(err)

	absFilePath := path.Join(dir, inputFilename)

	fmt.Println(absFilePath)
}
