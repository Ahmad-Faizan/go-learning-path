package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	pageTop = `<html>
				<head>
					<title>Soundex Algorithm</title>
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
)

func main() {
	http.HandleFunc("/", homeHandler)
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
				fmt.Fprintf(w, tableBody, name, len(name))
			}
			fmt.Fprint(w, tableBottom)
		}
	}
	fmt.Fprint(w, pageBottom)
}
