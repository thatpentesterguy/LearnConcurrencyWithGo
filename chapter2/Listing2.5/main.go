package main

import (
	"fmt"
	"runtime"
)

func sayHello() {

	fmt.Println("Hello from sayHello function")
}

func sayBye() {
	fmt.Println("Hello from sayBye function")
}

func main() {

	go sayHello()
	//go sayBye()
	// Wait for goroutines to finish
	runtime.Gosched()
	fmt.Println("Hello from main function")
}
