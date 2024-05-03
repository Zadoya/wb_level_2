package main

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

// выделил в отдельную функцию, чтобы протестить

func timeFunc() int {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Println(time)
	return 0
}

func main() {
	os.Exit(timeFunc())
}

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
