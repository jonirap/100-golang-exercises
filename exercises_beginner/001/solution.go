package main

import "fmt"

func main() {
	for i := 2000; i <= 3200; i++ {
		if i%7 == 0 && i%5 != 0 {
			fmt.Printf("%d, ", i)
		}
	}
}
