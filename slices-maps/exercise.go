package main

import "fmt"

func main() {
	result := uniqueInts([]int{9, 19, 5, 4, 3, 4, 4, 6, 3, 8, 9, 5, 2})
	fmt.Println(result)
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
