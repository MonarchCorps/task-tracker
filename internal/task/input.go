package task

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewScanner(os.Stdin)

func GetStringInput(prompt string) string {
	fmt.Print(prompt)
	reader.Scan()
	return reader.Text()
}

func GetIntInput(prompt string) int {
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
