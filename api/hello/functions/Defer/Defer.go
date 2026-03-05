//n Go language, defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns.
// In other words, defer function or method call arguments evaluate instantly, but they don't execute until the nearby functions returns.
// You can create a deferred method, or function, or anonymous function by using the defer keyword.

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
