package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("day3 part2")

	file, err := os.Open("day3/input.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	re := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)|do\(\)|don't\(\)`)
	//ops := make([]string, 0)
	sum := 0
	enableMul := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		matches := re.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				enableMul = true
			} else if match[0] == "don't()" {
				enableMul = false
			} else if match[1] != "" && match[2] != "" && enableMul {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				result := x * y
				sum += result
			} else if match[1] != "" && match[2] != "" {
				continue
			}
		}
	}
	fmt.Println(sum)
}
