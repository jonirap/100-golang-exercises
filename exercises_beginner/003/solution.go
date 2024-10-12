package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createIntegralMap(n int) map[int]int {
	integralMap := make(map[int]int)
	for i := 0; i <= n; i++ {
		integralMap[i] = i * i
	}
	return integralMap
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Please enter a number to integrate a map of (Enter nothing to exit): ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		if input == "" {
			break
		}
		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}
		integralMap := createIntegralMap(n)
		fmt.Println(integralMap)
	}
}
