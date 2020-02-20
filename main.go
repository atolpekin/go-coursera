package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, s string, b bool) error {

	files, err := ioutil.ReadDir("/Users/aktolpekin/Desktop/hw1_tree/testdata")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if range
		if  f.IsDir(){
			fmt.Println("└───" + f.Name())
		}
	}
	return err
}

