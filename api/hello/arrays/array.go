package main

import "fmt"

func main() {
	arr := [5]string{"Bob", "Alice", "John", "Eve", "Charlie"}

	fmt.Println("Elements Of The Array")

	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}
}
