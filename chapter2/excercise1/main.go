package main

import (
	"fmt"
	"os"
	"time"
)

// go run main.go txtfile1 txtfile2 txtfile3
func main() {

	// create file & write bubbles to it

	file, err := os.Create("txtfile1")
	if err != nil {
		fmt.Println("Cant create file", err)
	}
	defer file.Close()

	file2, err := os.Create("txtfile2")
	if err != nil {
		fmt.Println("Cant create file", err)
	}
	defer file2.Close()

	file3, err := os.Create("txtfile3")
	if err != nil {
		fmt.Println("Cant create file", err)
	}
	defer file3.Close()

	err = os.WriteFile("txtfile1", []byte("somebody broke my bubbles\n"), 0644)
	if err != nil {
		fmt.Println("Cant write to file1", err)
	}
	err = os.WriteFile("txtfile2", []byte("somebody stole my bubbles\n"), 0644)
	if err != nil {
		fmt.Println("Cant write to file", err)
	}
	err = os.WriteFile("txtfile3", []byte("somebody blew my bubbles\n"), 0644)
	if err != nil {
		fmt.Println("Cant write to file3", err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Please provide textfile names as go run main.go txtfile1 txtfile2 txtfile3")
	}

	files := os.Args[1:]
	fmt.Printf("Files to be read: %#v\n", files)

	for index, fileName := range files {

		switch index {

		case 0:

			go func() {
				data, err := os.ReadFile(fileName)
				if err != nil {
					panic(err)
				}
				fmt.Printf("Content of the file %s is: %s", fileName, data)
			}()

		case 1:

			go func() {
				data, err := os.ReadFile(fileName)
				if err != nil {
					panic(err)
				}
				fmt.Printf("Content of the file %s is : %s", fileName, data)
			}()

		case 2:

			go func() {
				data, err := os.ReadFile(fileName)
				if err != nil {
					panic(err)
				}
				fmt.Printf("Content of the file %s is : %s", fileName, data)
			}()

		default:
			fmt.Println("Please provide only 3 textfile names as go run main.go txtfile1 txtfile2 txtfile3")

		}

	}

	time.Sleep(2 * time.Second)
}
