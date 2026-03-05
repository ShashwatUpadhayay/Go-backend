package main

import "fmt"

func main() {

	for i, char := range "Go Programming" {
		fmt.Printf("The Index Number Of %c Is: %d\n", char, i)
	}
}
