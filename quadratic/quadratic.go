package main

import (
	"fmt"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

type equation struct {
	a  complex128
	b  complex128
	c  complex128
	x1 complex128
	x2 complex128
}

const (
	htmlPageTop     = `<html><head><title>QE Solver</title><style>.error{color:#FF0000;}</style></head>`
	htmlPageBody    = `<body><h2>Quadratic Equation Solver</h2><p>Fill the box with the values of coefficients in the equation.</p><form action="/" method="GET"><p><input type="text" name="a" size="1" placeholder="A">  X<sup>2</sup>  +  <input type="text" name="b" size="1" placeholder="B">  X  +  <input type="text" name="c" size="1" placeholder="C">  </p><p><input type="submit" value="Calculate"></p></form>`
	htmlPageBottom  = `</body></html>`
	htmlResult      = `<p>The roots of the above equation are <b>%v</b> and <b>%v</b>.</p>`
	htmlCustomError = `<p class="error">%s</p>`
)

func main() {
	http.HandleFunc("/", homePageGetHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("Unable to start server on the specified port ", err)
	}
}

func homePageGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlPageTop, htmlPageBody)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, htmlCustomError, err)
	} else {
		if values, errInParsing, ok := processForm(r); ok {
			var eqn equation
			eqn.a, eqn.b, eqn.c = complex(values[0], 0), complex(values[1], 0), complex(values[2], 0)
			eqn.x1, eqn.x2 = calculateRoots(eqn)
			fmt.Fprint(w, formatResult(eqn))
		} else if errInParsing != "" {
			fmt.Fprintf(w, htmlCustomError, errInParsing)
		}
	}
	fmt.Fprint(w, htmlPageBottom)
}

func processForm(r *http.Request) ([3]float64, string, bool) {
	var values [3]float64
	count := 0
	for i, param := range []string{"a", "b", "c"} {
		if strValues, found := r.Form[param]; found && len(strValues) == 1 {
			if strValues[0] != "" {
				if x, err := strconv.ParseFloat(strValues[0], 64); err != nil {
					return values, strValues[0] + " is an invalid number.", false
				} else {
					values[i] = x
				}
			} else {
				values[i] = 0
			}
			count++
		}
	}
	if count != 3 {
		return values, "", false
	}
	if values[0] == 0 {
		return values, "Coefficient of x<sup>2</sup> can't be 0", false
	}
	return values, "", true
}

func formatResult(eqn equation) string {
	return fmt.Sprintf(htmlResult, eqn.x1, eqn.x2)
}

func calculateRoots(eqn equation) (complex128, complex128) {
	d := (eqn.b * eqn.b) - (4 * eqn.a * eqn.c)
	eqn.x1 = (-eqn.b + cmplx.Sqrt(d)) / (2 * eqn.a)
	eqn.x2 = (-eqn.b - cmplx.Sqrt(d)) / (2 * eqn.a)
	return eqn.x1, eqn.x2
}
