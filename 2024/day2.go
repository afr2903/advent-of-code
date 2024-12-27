package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_safe(numbers []string) bool {
	var past_number int
	var increasing bool
	for i := 0; i < len(numbers); i++ {
		number, err := strconv.Atoi(numbers[i])
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return false
		}
		if i == 0 {
			past_number = number
			continue
		}
		if i == 1 {
			if number > past_number {
				increasing = true
			} else {
				increasing = false
			}
		}
		if increasing {
			if number <= past_number {
				return false
			}
		} else {
			if number >= past_number {
				return false
			}
		}
		if abs(number, past_number) > 3 {
			return false
		}
		past_number = number
	}
	return true
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	// First part: Find the number of safe reports
	safe_reports := 0
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if is_safe(numbers) {
			safe_reports++
		} else {
			for i := 0; i < len(numbers); i++ {
				new_numbers := make([]string, 0)
				new_numbers = append(new_numbers, numbers[:i]...)
				new_numbers = append(new_numbers, numbers[i+1:]...)
				if is_safe(new_numbers) {
					safe_reports++
					break
				}
			}
		}
	}
	fmt.Println("Number of safe reports: ", safe_reports)
}

func abs(i int, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
