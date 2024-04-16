package main

/*
import (
	"fmt"
)

func largestRectangleArea(heights []int) int {
	ans, stack := 0, make([]int, 0)
	for i := 0; i <= len(heights); i++ {
		for len(stack) > 0 && (i == len(heights) || heights[i] <= heights[stack[len(stack)-1]]) {
			removedHeight := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			ans = max(removedHeight * width, ans)
		}
		stack = append(stack, i)
	}
	return ans
}

func maximalRectangle(matrix [][]byte) int {
	maxSquare := 0
	height := make([]int, len(matrix[0]))

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != 0 {
				height[j]++
			} else {
				height[j] = 0
			}
		}

		maxSquare = max(maxSquare, largestRectangleArea(height))
	}
	return maxSquare
}

func main() {
	//matrix := [][]byte{{0,1}, {0,1}}
	matrix := [][]byte{{1, 0, 1, 0, 0, 0}, {1, 0, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 0, 0, 1, 0, 0}}

	fmt.Println(maximalRectangle(matrix))
}
*/