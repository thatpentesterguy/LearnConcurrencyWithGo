package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func countWords(url string, frequency map[string]int) {

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	lower_case_words := strings.ToLower(string(body))

	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)

	for _, words := range wordRegex.FindAllString(lower_case_words, -1) {
		frequency[words]++
	}

	fmt.Println("Completed:", url)

}

func main() {

	var frequency = make(map[string]int)

	for i := 1000; i <= 1030; i++ {

		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWords(url, frequency)
	}
	time.Sleep(10 * time.Second) // wait for goroutines to finish
	for i, v := range frequency {
		fmt.Printf("%s-->%d ", i, v)
	}

}
