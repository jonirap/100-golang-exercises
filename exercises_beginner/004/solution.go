package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sequence of comma-separated numbers: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strNumbers := strings.Split(input, ",")
	var numbers []int
	for _, str := range strNumbers {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	fmt.Println(numbers)
}
