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

	resultMatrix := make2D([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 6)
	fmt.Println(resultMatrix)
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
