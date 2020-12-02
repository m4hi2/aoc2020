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
		parts := strings.Split(scanner.Text(), ":")
		rule, password := parts[0], parts[1]

		password = strings.TrimSpace(password)
		if passwordChecker(password, rule) {
			counter++
		}
	}
	fmt.Println(counter)
}

func passwordChecker(password string, rule string) bool {
	ruleParts := strings.Split(rule, " ")
	pos, key := ruleParts[0], ruleParts[1]
	posParts := strings.Split(pos, "-")
	firstPos, _ := strconv.Atoi(posParts[0])
	secondPos, _ := strconv.Atoi(posParts[1])
	firstPos = firstPos - 1
	secondPos = secondPos - 1

	flag := false

	if password[firstPos] == key[0] {
		flag = true
	}

	if password[secondPos] == key[0] {
		if flag {
			return false
		}
		flag = true

	}
	return flag
}
