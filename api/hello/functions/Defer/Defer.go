package main

import "fmt"

func add(a, b int) int {
	res := a + b
	fmt.Println("The Result Is: ", res)
	return 0
}

func main() {

	fmt.Println("Start")
	defer fmt.Println("End")
	defer fmt.Println("Output Presented In LIFO Order")
	defer add(5, 10)
	defer add(20, 30)
}
