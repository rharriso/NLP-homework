package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("\n Usage: countFreq path/to/inputFile \n")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err.Error())
	}

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		members := strings.SplitAfter(scanner.Text(), " ")
		if len(members) == 2 {
			fmt.Printf("%s \t aahhahah \t %s\n", members[0], members[1])
		}
	}
}
