package main

import (
	"os"
)

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
}
