package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	pos := 0
	tree := 0
	file, _ := os.Open("./input_file.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		pos += 3
		if string(line[pos%len(line)]) == "#" {
			tree++
		}
	}
	fmt.Println(tree)
}
