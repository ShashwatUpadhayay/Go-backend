package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println("Total:", total)
	return 0
}

func greet(prefix string, nums ...int) {
	fmt.Println(prefix)
	for _, n := range nums {
		fmt.Println(n)
	}
}
func main() {
	greet("Numbers:", 1, 2, 3, 4, 5)
	greet("More numbers:", 10, 20)
	greet("No numbers:")
	sum(1, 2, 3, 4, 5)
	sum(10, 20)
	sum()
}
