package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("file", "", "the file to parse the separator from")
	flag.Parse()

	if *fileName == "" || len(os.Args) == 1 {
		fmt.Print("No file specified")
		os.Exit(1)
	}

	reader, err := os.Open(*fileName)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	lines, n := readLines(reader)
	separator := findSeparator(lines, n)
	printReport(separator)
}

func readLines(r io.Reader) ([]string, int) {
	reader := bufio.NewReader(r)
	maxLines := 5
	var lines []string
	for i := 0; i < maxLines; i++ {
		line, _ := reader.ReadString('\n')
		lines = append(lines, line)
	}
	return lines, maxLines
}

func findSeparator(lines []string, n int) string {
	separators := []string{"\t", "*", "|", ","}
	count := make([][]int, len(separators))

	for sepIndex := range separators {
		count[sepIndex] = make([]int, n)
		for lineIndex, line := range lines {
			count[sepIndex][lineIndex] = strings.Count(line, separators[sepIndex])
		}
	}
	return checkSeparator(count, separators, n)
}

func checkSeparator(count [][]int, separators []string, n int) string {
	for i := range separators {
		firstVal := count[i][0]
		ok := true
		for j := 1; j < n; j++ {
			if count[i][j] != firstVal {
				ok = false
				break
			}
		}
		if firstVal > 0 && ok {
			return separators[i]
		}
	}
	return "whitespace"
}
func printReport(sep string) {
	fmt.Printf("The file is %s separated.", sep)
	return
}
