package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var leftIDs, rightIDs []int

	similarityScore := 0

	rightListNumberCount := make(map[int]int)

	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Fields(line)
		if len(cols) == 2 {
			left, _ := strconv.Atoi(cols[0])
			right, _ := strconv.Atoi(cols[1])
			leftIDs = append(leftIDs, left)
			rightIDs = append(rightIDs, right)
		}
	}

	for _, val := range rightIDs {
		rightListNumberCount[val]++
	}

	for _, leftVal := range leftIDs {
		similarityScore += leftVal * rightListNumberCount[leftVal]
	}

	fmt.Println(similarityScore)

}
