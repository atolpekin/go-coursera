package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
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

func dirTree(out io.Writer, path string, printFiles bool) error {
	recursivePrint("", out, path, printFiles)
	return nil
}

func recursivePrint(indent string, output io.Writer, currDir string, printFiles bool) {
	fileObj, _ := os.Open(currDir)
	defer fileObj.Close()

	fileName := fileObj.Name()
	files, _ := ioutil.ReadDir(fileName)

	var filesMap map[string]os.FileInfo = map[string]os.FileInfo{}
	var unSortedFiles []string = []string{}
	for _, file := range files {
		unSortedFiles = append(unSortedFiles, file.Name())
		filesMap[file.Name()] = file
	}
	sort.Strings(unSortedFiles)
	var sortedFiles []os.FileInfo = []os.FileInfo{}
	for _, stringName := range unSortedFiles {
		sortedFiles = append(sortedFiles, filesMap[stringName])
	}
	files = sortedFiles
	var newFileList []os.FileInfo = []os.FileInfo{}
	var length int
	if !printFiles {
		for _, file := range files {
			if file.IsDir() {
				newFileList = append(newFileList, file)
			}
		}
		files = newFileList
	}
	length = len(files)
	for i, file := range files {
		if file.IsDir() {
			var stringPrepender string
			if length > i+1 {
				fmt.Fprintf(output, indent+"├───"+"%s\n", file.Name())
				stringPrepender = indent + "│\t"
			} else {
				fmt.Fprintf(output, indent+"└───"+"%s\n", file.Name())
				stringPrepender = indent + "\t"
			}
			newDir := filepath.Join(currDir, file.Name())
			recursivePrint(stringPrepender, output, newDir, printFiles)
		} else if printFiles {
			if file.Size() > 0 {
				if length > i+1 {
					fmt.Fprintf(output, indent+"├───%s (%vb)\n", file.Name(), file.Size())
				} else {
					fmt.Fprintf(output, indent+"└───%s (%vb)\n", file.Name(), file.Size())
				}
			} else {
				if length > i+1 {
					fmt.Fprintf(output, indent+"├───%s (empty)\n", file.Name())
				} else {
					fmt.Fprintf(output, indent+"└───%s (empty)\n", file.Name())
				}
			}
		}
	}
}
