package main

import (
	"bufio"
	"fmt"
	"os"
)

func countXmas(grid [][]rune) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	directions := [][2]int{
		{0, 1},   // right
		{1, 1},   // down-right
		{1, 0},   // down
		{1, -1},  // down-left
		{0, -1},  // left
		{-1, -1}, // up-left
		{-1, 0},  // up
		{-1, 1},  // up-right
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'X' {
				for _, dir := range directions {
					if checkDirection(grid, i, j, dir[0], dir[1]) {
						count++
					}
				}
			}
		}
	}
	return count
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func checkDirection(grid [][]rune, x, y, m, n int) bool {
	rows, cols := len(grid), len(grid[0])
	target := []rune{'X', 'M', 'A', 'S'}

	for i := 0; i < len(target); i++ {
		newX, newY := x+i*m, y+i*n
		if !isValid(newX, newY, rows, cols) || grid[newX][newY] != target[i] {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	result := countXmas(grid)
	fmt.Println(result)
}
