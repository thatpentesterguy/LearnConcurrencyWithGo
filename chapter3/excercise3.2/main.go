package main

import (
	"fmt"
	"time"
)

func addNextNumber(nextNum *[101]int) {
	i := 0
	for nextNum[i] != 0 { // Find the next available index
		i++
	}
	nextNum[i] = nextNum[i-1] + 1
}

func main() {
	nextNum := [101]int{1}

	for i := 0; i < 100; i++ { // calling 100 times to fill the array
		go addNextNumber(&nextNum)

	}

	for nextNum[100] == 0 {
		println("Waiting for computation to finish...")
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println(nextNum)
}
