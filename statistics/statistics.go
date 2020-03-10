package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

var (
	htmlPageHeader  = `<!DOCTYPE HTML><html><head><style>.error{color:#FF0000;}</style></head><title>Statistics</title><body><h3>Statistics</h3><p>Computes basic statistics for a given list of numbers</p>`
	htmlPageFooter  = `</body></html>`
	htmlFormBody    = `<form action="/" method="POST"><label for="numbers">Numbers (comma or space-separated):</label><br /><input type="text" name="numbers" size="30"><br /><input type="submit" value="Calculate"></form>`
	htmlTableData   = `<table border="1"><tr><th colspan="2">Results</th></tr><tr><td>Numbers</td><td>%v</td></tr><tr><td>Count</td><td>%d</td></tr><tr><td>Mean</td><td>%f</td></tr><tr><td>Median</td><td>%f</td></tr></table>`
	htmlCustomError = `<p class="error">%s</p>`
)

func main() {
	http.HandleFunc("/", homePageGetHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("Server did not start on the specified port", err)
	}
}

func homePageGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlPageHeader, htmlFormBody)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, htmlCustomError, err)
	} else {
		if numbers, errBody, ok := processForm(r); ok {
			stats := getStats(numbers)
			fmt.Fprintf(w, formatStats(stats))
		} else if errBody != "" {
			fmt.Fprintf(w, htmlCustomError, errBody)
		}
	}
	fmt.Fprintf(w, htmlPageFooter)
}

func processForm(r *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, ok := r.Form["numbers"]; ok && len(slice) > 0 {
		filteredSlice := strings.ReplaceAll(slice[0], ",", " ")
		for _, field := range strings.Fields(filteredSlice) {
			if num, err := strconv.ParseFloat(field, 64); err == nil {
				numbers = append(numbers, num)
			} else {
				return numbers, "&lt;" + field + "&gt; is invalid", false
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false
	}
	return numbers, "", true
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(htmlTableData, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(stats.numbers) / float64(len(stats.numbers))
	stats.median = median(stats.numbers)
	return stats
}

func sum(numbers []float64) (total float64) {
	for _, i := range numbers {
		total += i
	}
	return total
}

func median(numbers []float64) (result float64) {
	middle := len(numbers) / 2
	result = numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}
