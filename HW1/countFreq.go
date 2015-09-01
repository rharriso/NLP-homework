package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var wordCount map[string]int

func main() {
	if len(os.Args) != 2 {
		fmt.Println("\n Usage: countFreq path/to/inputFile \n")
		return
	}

	// open the passed file
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err.Error())
	}

	wordCount = make(map[string]int)

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		members := strings.SplitAfter(scanner.Text(), " ")
		if len(members) == 2 {
			wordCount[members[0]]++
		}
	}

	// print word counts
	for word, count := range wordCount {
		fmt.Println(word, count)
	}
}
