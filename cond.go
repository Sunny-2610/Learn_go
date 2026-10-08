package main

import "fmt"

func main() {
	x := 10

	if x > 0 {
		fmt.Println("Positive")
	} else if x == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Negative")
	}
}