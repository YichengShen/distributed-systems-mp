package utils

import (
	"bufio"
	"os"
)

// helper function to take user inputs
func TakeUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}