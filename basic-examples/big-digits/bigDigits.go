package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string{
	{
		"   000   ",
		"  0   0  ",
		" 0     0 ",
		" 0     0 ",
		" 0     0 ",
		"  0   0  ",
		"   000   ",
	},
	{
		"   1  ",
		" 111  ",
		"   1  ",
		"   1  ",
		"   1  ",
		"   1  ",
		" 1111 ",
	},
	{
		"   2222  ",
		" 22   22 ",
		"      22 ",
		"    22   ",
		"   22    ",
		" 22      ",
		" 2222222 ",
	},
	{
		"   3333   ",
		" 3    33  ",
		"       33 ",
		"   3333   ",
		"      33  ",
		" 3     33 ",
		"   3333   ",
	},
	{
		"       4   ",
		"     4 4   ",
		"   4   4   ",
		" 4     4   ",
		" 444444444 ",
		"      4    ",
		"       4   ",
	},
	{
		" 5555555  ",
		" 5        ",
		" 5        ",
		"  55555   ",
		"       55 ",
		" 5    55  ",
		"  5555    ",
	},
	{
		"   6666  ",
		"  6      ",
		" 6       ",
		" 6666666 ",
		" 6     6 ",
		"  6   6  ",
		"   666   ",
	},
	{
		" 77777777 ",
		"      77  ",
		"     77   ",
		"    77    ",
		"   77     ",
		"  77      ",
		" 77       ",
	},
	{
		"  88888  ",
		" 8     8 ",
		" 88   88 ",
		"  88888  ",
		" 88   88 ",
		" 8     8 ",
		"  88888  ",
	},
	{
		"  999999  ",
		" 9     99 ",
		"  9    99 ",
		"   999999 ",
		"       99 ",
		"       99 ",
		"       99 ",
	},
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] {
		line := ""
		for col := range stringOfDigits {
			digit := stringOfDigits[col] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + "  "
			} else {
				log.Fatal("invalid number entered")
			}
		}
		fmt.Println(line)
	}
}
