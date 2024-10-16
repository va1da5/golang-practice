package main

// go run main.go "\S+?e\s" the-road-not-taken.txt

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: egrep <pattern> <text-file>")
		os.Exit(0)
	}

	pattern, err := regexp.Compile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filename := os.Args[2]

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if pattern.Match([]byte(line)) {
			match := pattern.Find([]byte(line))
			fmt.Printf("%s\n", strings.ReplaceAll(line, string(match), Red+string(match)+Reset))
		}

	}

}
