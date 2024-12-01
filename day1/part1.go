package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	var total int

	var leftIDs, rightIDs []int

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

	sort.Ints(leftIDs)
	sort.Ints(rightIDs)

	for idx := range leftIDs {
		if leftIDs[idx] > rightIDs[idx] {
			total += leftIDs[idx] - rightIDs[idx]
		} else {
			total += rightIDs[idx] - leftIDs[idx]
		}
	}

	fmt.Println(total)
}
