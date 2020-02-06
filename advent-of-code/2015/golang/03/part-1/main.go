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

func getFileData(filename string) []byte {
	wd, err := os.Getwd()

	handleErr(err)

	absFilePath := path.Join(wd, string(filename))
	data, err := ioutil.ReadFile(absFilePath)

	handleErr(err)

	return data
}

func main() {
	inputFile := os.Args[1:][0]
	data := getFileData(inputFile)

	fmt.Printf("%s\n", data)
}
