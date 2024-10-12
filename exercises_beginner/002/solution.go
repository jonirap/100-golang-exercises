package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter number to factorize (enter nothing to quit): ")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			break
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}
		fmt.Printf("Factorial of %d is %d\n", number, factorial(number))
	}
}
