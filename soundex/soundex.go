package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	pageTop = `<html>
				<head>
					<title>Soundex Algorithm</title>
					<style>.error{ color : #FF0000} .pass{ color : #00FF00} .fail{ color : #FF0909}</style>
				</head>
				<body>
					<h1>Soundex</h1>
					<p>Compute Soundex codes for a list of names</p>`
	formBody = `<form action="/" method="GET">
					<p>Names (Comma or space separated)</p>
					<input type="text" name="names" width="300px">
					<input type="submit" value="Compute">
				</form>`
	tableHead = `<table border="1">
					<th>Name</th><th>Score</th>`
	tableBody   = `<tr><td>%v</td><td>%v</td></tr>`
	tableBottom = `</table>`
	pageBottom  = `</body>
					</html>`
	htmlError = `<p class="error">%s</p>`
	testTable = `<table border = "1">
				<th>Name</th><th>Calculated Score</th><th>Expected Score</th><th>Result</th>`
	testTableData = `<tr><td>%v</td><td>%v</td><td>%v</td><td>%v</td></tr>`
)

var digitForLetter = []rune{
	//A  B  C  D  E  F  G  H  I  J  K  L  M
	0, 1, 2, 3, 0, 1, 2, 0, 0, 2, 2, 4, 5,
	//N  O  P  Q  R  S  T  U  V  W  X  Y  Z
	5, 0, 1, 2, 6, 2, 3, 0, 1, 0, 2, 0, 2}

var testCases map[string]string

func main() {
	http.HandleFunc("/", homeHandler)
	var ok bool
	if testCases, ok = loadTestsFromTXT("soundex-test-data.txt"); ok {
		http.HandleFunc("/test", testCaseHandler)
	}
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal("Server did not start on the specifed port.\n", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, pageTop, formBody)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, htmlError, err)
	} else {
		if rawValues, ok := r.Form["names"]; ok && len(rawValues) > 0 {
			filteredValues := strings.ReplaceAll(rawValues[0], ",", " ")
			fmt.Fprint(w, tableHead)
			for _, name := range strings.Fields(filteredValues) {
				fmt.Fprintf(w, tableBody, name, calcScore(name))
			}
			fmt.Fprint(w, tableBottom)
		}
	}
	fmt.Fprint(w, pageBottom)
}

func testCaseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, pageTop, testTable)
	var buf bytes.Buffer
	for name, score := range testCases {
		soundex := calcScore(name)
		var testResult string
		if soundex == score {
			testResult = `<span class="pass">PASS</span>`
		} else {
			testResult = `<span class="fail">FAIL</span>`
		}
		buf.WriteString(formatData(name, soundex, score, testResult))
	}
	fmt.Fprint(w, buf.String(), tableBottom, pageBottom)
}

func formatData(name, calculatedScore, expectedScore, testResult string) string {
	return fmt.Sprintf(testTableData, name, calculatedScore, expectedScore, testResult)
}

func calcScore(name string) string {
	name = strings.ToUpper(name)
	chars := []rune(name)
	var codes []rune
	codes = append(codes, chars[0])
	for i := 1; i < len(chars); i++ {
		char := chars[i]
		if i > 0 && char == chars[i-1] {
			continue
		}
		if index := char - 'A'; 0 <= index &&
			index < int32(len(digitForLetter)) &&
			digitForLetter[index] != 0 {
			codes = append(codes, '0'+digitForLetter[index])
		}
	}
	for len(codes) < 4 {
		codes = append(codes, '0')
	}
	return string(codes[:4])
}

func loadTestsFromTXT(path string) (map[string]string, bool) {
	testCases := make(map[string]string)
	if rawData, err := ioutil.ReadFile(path); err != nil {
		log.Print(err)
		return testCases, false
	} else {
		for _, line := range strings.Split(string(rawData), "\n") {
			values := strings.Fields(line)
			if len(values) == 2 {
				testCases[values[1]] = values[0]
			}
		}
		return testCases, true
	}
}
