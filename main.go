package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func countLines(r io.Reader) (int, int, int) {
	lineScanner := bufio.NewScanner(r)
	var lineCount, wordCount, charCount int
	re := regexp.MustCompile(`[\S]+`)

	for lineScanner.Scan() {
		lineCount++
		lineText := lineScanner.Text()
		wordCount += len(re.FindAllString(lineText, -1))
		//fmt.Println(len(strings.Split(lineText, " ")))
		//wordCount +=
		charCount += len(lineText)
	}
	return lineCount, wordCount, charCount
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
