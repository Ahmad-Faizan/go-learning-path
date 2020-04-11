package main

import "fmt"

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
