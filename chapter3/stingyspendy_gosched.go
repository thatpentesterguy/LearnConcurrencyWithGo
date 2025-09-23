package main

import (
	"fmt"
	"time"
	"runtime"
)

func stingy(money *int) {

	for i :=0; i < 1000000; i++{
		*money += 10
		runtime.Gosched() // yield the processor to allow other goroutines to run
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int) {
	for i :=0; i < 1000000; i++{
		*money -= 10
		runtime.Gosched() // yield the processor to allow other goroutines to run
	}
	fmt.Println("Spendy Done")
}

func main() {

	money := 100 
	
	
	go stingy(&money) // address of money 

	go spendy(&money) // address of money
	
	time.Sleep(2 * time.Second)
	fmt.Println("Money in Bank Account:", money)
}

