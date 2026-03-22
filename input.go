package main

import (
	"fmt"
	"strconv"
)

func getStringInput(prompt string) string {
	fmt.Print(prompt)
	reader.Scan()
	return reader.Text()
}

func getIntInput(prompt string) int {
	for {
		fmt.Print(prompt)
		reader.Scan()
		input := reader.Text()

		num, err := strconv.Atoi(input)
		if err == nil {
			return num
		}

		fmt.Println("Invalid number. Try again.")
	}
}
