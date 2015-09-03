package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var wordCount map[string]int
var biTagCount map[string]int
var triTagCount map[string]int

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
	biTagCount = make(map[string]int)
	triTagCount = make(map[string]int)

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var tags []string

	for scanner.Scan() {
		members := strings.SplitAfter(scanner.Text(), " ")

		if tags == nil || len(members) < 2 {
			tags = []string{"*", "*", "*"}
		}

		if len(members) == 2 {
			wordCount[members[0]]++
			tags = append(tags[1:], members[1])
			biTagCount[strings.Join(tags[1:], " ")]++
			triTagCount[strings.Join(tags, " ")]++
		}
	}

	// print word counts
	for word, count := range wordCount {
		fmt.Println(word, count)
	}

	fmt.Println("\nBi Tag Count:")
	for tag, count := range biTagCount {
		fmt.Println(tag, count)
	}

	fmt.Println("\nTri Tag Count:")
	for tag, count := range triTagCount {
		fmt.Println(tag, count)
	}
}
