package main

import (
	"net/http"
	"io"
	"strings"
	"fmt"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int) {

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}

	body, _ := io.ReadAll(resp.Body) // []byte

	for _, b := range body {

		c := strings.ToLower(string(b))

		cIndex := strings.Index(allLetters, c)

		if cIndex >= 0 {
			frequency[cIndex] += 1
		}


	}
	fmt.Println("Completed:", url)


}

func main() {

	var frequency = make([]int, 26) // can be len(allLetters)

	for i := 1000 ; i <= 1030; i++ {

		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		countLetters(url, frequency)
	}
	for i, v := range allLetters {
		fmt.Printf("%c-%d ", v, frequency[i])
	}

}