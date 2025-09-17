package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// grep function to search a word in a file

func grepString(filename string, searchstr string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if strings.Contains(string(content), searchstr) {
		fmt.Printf("Found the word %s in file %s\n", searchstr, filename)
	} else {
		fmt.Printf("Did not find the word %s in file %s\n", searchstr, filename)
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
	if searchstr == "" {
		fmt.Println("Please provide a word to be searched as go run main.go <word> <file1> <file2> <file3>")
		return
	}
	filenames := os.Args[2:] // list of files to be searched
	if len(filenames) == 0 {
		fmt.Println("Please provide textfile names as go run main.go <word> <file1> <file2> <file3>")
		return
	}

	for _, filename := range filenames {
		go grepString(filename, searchstr)
	}

	time.Sleep(2 * time.Second)

}
