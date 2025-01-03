package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(i int, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	// First part: Find the total distance between the two lists
	var left []int
	var right []int
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		left_number, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return
		}
		right_number, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return
		}
		left = append(left, left_number)
		right = append(right, right_number)
	}
	sort.Ints(left)
	sort.Ints(right)
	total_distance := 0
	for i := 0; i < len(left); i++ {
		total_distance += abs(left[i], right[i])
	}
	fmt.Println("Total distance: ", total_distance)

	// Second part: Find the similarity score between the two lists
	var right_map map[int]int
	right_map = make(map[int]int)
	for i := 0; i < len(right); i++ {
		if _, ok := right_map[right[i]]; ok {
			right_map[right[i]]++
		} else {
			right_map[right[i]] = 1
		}
	}
	similarity_score := 0
	for i := 0; i < len(left); i++ {
		if _, ok := right_map[left[i]]; ok {
			similarity_score += left[i] * right_map[left[i]]
		}
	}
	fmt.Println("Similarity score: ", similarity_score)
}
