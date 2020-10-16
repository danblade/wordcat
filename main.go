package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func countLines(r io.Reader) (int, int, int) {
	lineScanner := bufio.NewScanner(r)
	var lineCount, wCount, charCount int

	for lineScanner.Scan() {
		lineCount++
		lineText := lineScanner.Text()
		fmt.Printf("%#v\n", strings.Fields(lineText))
		wCount += len(strings.Fields(lineText))
		charCount += len(lineText)
	}
	return lineCount, wCount, charCount
}

func main() { //My WordCat
	var r io.Reader

	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		r = file
	} else if len(os.Args) == 1 {
		r = os.Stdin
	} else {
		fmt.Println("Usage: wordcat  <file to examine>")
		os.Exit(0)
	}
	fmt.Println(countLines(r))
}
