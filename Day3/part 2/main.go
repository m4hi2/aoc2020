package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := "./input_file.txt"
	final := treeFinder(1, 1, file) * treeFinder(3, 1, file) * treeFinder(5, 1, file) * treeFinder(7, 1, file) * treeFinder(1, 2, file)
	fmt.Println(final)

}

func treeFinder(right int, down int, fileName string) int {
	// let's get the file working first
	pos := 0
	tree := 0
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for true {
		pos += right
		for i := 0; i < down; i++ {
			scanner.Scan()
		}
		if scanner.Text() == "" {
			break
		}
		if string(scanner.Text()[pos%len(scanner.Text())]) == "#" {
			tree++
		}
	}
	return tree
}
