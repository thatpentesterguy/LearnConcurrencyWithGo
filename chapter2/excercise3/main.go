package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// grep function to search a word in a directory

func grepPath(path string, fileinfo os.DirEntry, searchstr string) {

	fullpath := filepath.Join(path, fileinfo.Name())

	if !fileinfo.IsDir() {
		content, err := os.ReadFile(fullpath)

		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(string(content), searchstr) {
			fmt.Printf("Found the word %s in file %s\n at path %s\n", searchstr, fileinfo.Name(), path)
		} else {
			fmt.Printf("Did not find the word %s in file %s\n at path %s\n", searchstr, fileinfo.Name(), path)
		}
	}

}

func main() {

	err := os.WriteFile("txtfile1", []byte("somebody broke my bubbles\n"), 0644)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("txtfile2", []byte("i can blow bubbles in a ballon\n"), 0644)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("txtfile3", []byte("i am a go programmer\n"), 0644)

	if err != nil {
		panic(err)
	}

	searchstr := os.Args[1] // word to be searched

	path := os.Args[2] // path to be searched

	file, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, fileinfo := range file {

		go grepPath(path, fileinfo, searchstr)
	}

	time.Sleep(2 * time.Second)

}
