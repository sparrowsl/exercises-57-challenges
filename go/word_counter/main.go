package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse() // parse all flags provided by user

	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wordCount := 0

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}
