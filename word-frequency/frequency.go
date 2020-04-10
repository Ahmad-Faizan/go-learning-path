package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	fileName := flag.String("file", "", "the filename containing the text for frequency checking.")
	flag.Parse()

	if *fileName == "" || len(os.Args) == 1 {
		fmt.Print("File not specified.")
		os.Exit(1)
	}

	rawData, err := os.Open(*fileName)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	freq, words := getFrequency(rawData)
	invFreq, keys := getInverseFrequency(freq)

	for _, word := range words {
		fmt.Printf("%-10s: \t %d\n", word, freq[word])
	}

	for _, key := range keys {
		fmt.Printf("Count: %d %v\n", key, invFreq[key])
	}
}

func getFrequency(data io.Reader) (map[string]int, []string) {
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanWords)

	freq := make(map[string]int)
	var words []string
	for scanner.Scan() {
		word := strings.TrimSuffix(strings.ToLower(scanner.Text()), ".")
		_, present := freq[word]
		if present {
			freq[word]++
		} else {
			freq[word] = 1
			words = append(words, word)
		}
	}
	sort.Strings(words)
	return freq, words
}

func getInverseFrequency(freq map[string]int) (map[int][]string, []int) {
	invFreq := make(map[int][]string)
	var keys []int
	for word, count := range freq {
		invFreq[count] = append(invFreq[count], word)
	}

	for key := range invFreq {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	return invFreq, keys
}
