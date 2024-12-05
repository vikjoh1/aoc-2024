package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	safeCount := 0

	file, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		var numbers []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Printf("error converting %s to int: %v\n", field, err)
				continue
			}
			numbers = append(numbers, num)
		}

		if isSafeWithDampener(numbers) {
			safeCount++
		}
	}

	fmt.Println(safeCount)

}

func isSafeWithDampener(numbers []int) bool {
	if isSafe(numbers) {
		return true
	}

	for i := 0; i < len(numbers); i++ {
		dampened := make([]int, 0, len(numbers)-1)
		dampened = append(dampened, numbers[:i]...)
		dampened = append(dampened, numbers[i+1:]...)

		if isSafe(dampened) {
			return true
		}
	}

	return false
}

func isSafe(numbers []int) bool {
	if len(numbers) < 2 {
		return true
	}

	firstDiff := numbers[1] - numbers[0]
	if firstDiff == 0 || abs(firstDiff) > 3 {
		return false
	}

	increasing := firstDiff > 0
	for i := 2; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]

		if (increasing && diff <= 0) || (!increasing && diff >= 0) || abs(diff) > 3 || abs(diff) == 0 {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
