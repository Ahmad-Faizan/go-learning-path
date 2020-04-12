package main

import (
	"fmt"
	"strings"
)

type groupMap map[string]string

func main() {
	result := uniqueInts([]int{9, 19, 5, 4, 3, 4, 4, 6, 3, 8, 9, 5, 2})
	fmt.Println(result)

	res := flattenMatrix([][]int{
		{1, 2, 3, 4},
		{5, 6, 7},
		{8, 9, 10, 11, 12, 13},
		{14, 15, 16, 17, 18},
		{19, 20}},
	)
	fmt.Println(res)

	resultMatrix := make2D([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 6)
	fmt.Println(resultMatrix)

	dataINI := []string{
		"; Cut down copy of Mozilla application.ini file",
		"",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0",
		"EnableExtensionManager=1",
	}
	mapINI := parseINI(dataINI)
	prettyPrintINI(mapINI)
}

func uniqueInts(duplicates []int) (unique []int) {
	duplicateMap := map[int]bool{}
	for _, val := range duplicates {
		if _, exists := duplicateMap[val]; exists {
			continue
		} else {
			duplicateMap[val] = true
			unique = append(unique, val)
		}
	}
	return unique
}

func flattenMatrix(matrix [][]int) (slice []int) {
	for _, internalSlice := range matrix {
		for _, x := range internalSlice {
			slice = append(slice, x)
		}
	}
	return slice
}

func make2D(numbers []int, columns int) [][]int {
	getRows := func(totalSize int, columns int) int {
		if totalSize%columns != 0 {
			return (1 + totalSize/columns)
		}
		return totalSize / columns
	}

	matrix := make([][]int, getRows(len(numbers), columns))

	for i, x := range numbers {
		row := i / columns
		col := i % columns
		if matrix[row] == nil {
			matrix[row] = make([]int, columns)
		}
		matrix[row][col] = x
	}
	return matrix
}

func parseINI(data []string) map[string]groupMap {
	iniMap := make(map[string]groupMap)
	g := make(groupMap)
	var groupName string
	for _, line := range data {
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, ";") {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			line = line[1 : len(line)-1]
			groupName = line
			g = map[string]string{}
		} else if sepIndex := strings.Index(line, "="); sepIndex != -1 {
			key, value := line[:sepIndex], line[sepIndex+1:]
			g[key] = value
			iniMap[groupName] = g
		}
	}
	return iniMap
}

func prettyPrintINI(iniMap map[string]groupMap) {
	fmt.Print("map[")
	size := 0
	for name, g := range iniMap {
		fmt.Printf("%s%s: map[", strings.Repeat(" ", size), name)
		for gKey, gValue := range g {
			fmt.Printf("%s: %s\n%s", gKey, gValue, strings.Repeat(" ", len("map[")+len(name+": map[")))
		}
		fmt.Print("\b]\n")
		size = len("map[")
	}
	fmt.Print("]")
}
