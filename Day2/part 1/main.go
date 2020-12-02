package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input_data.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		parts := strings.Split(scanner.Text(), ":")
		rule, password := parts[0], parts[1]

		if passwordChecker(password, rule) {
			counter++
		}
	}
	fmt.Println(counter)
}

func passwordChecker(password string, rule string) bool {
	ruleParts := strings.Split(rule, " ")
	limit, key := ruleParts[0], ruleParts[1]
	limitParts := strings.Split(limit, "-")
	min, _ := strconv.Atoi(limitParts[0])
	max, _ := strconv.Atoi(limitParts[1])
	keyCount := strings.Count(password, key)
	if keyCount >= min && keyCount <= max {
		return true
	}
	return false

}
