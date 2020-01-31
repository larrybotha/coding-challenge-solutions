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

func main() {
	data := getInputData()

	fmt.Printf("file data:\n %s", data)
}
