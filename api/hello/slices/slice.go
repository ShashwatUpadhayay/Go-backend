package main

import (
	"fmt"
	"sort"
)

func main() {
	array := [5]string{"Bob", "Alice", "John", "Eve", "Charlie"}

	slice := array[2:4]
	slice = append(slice, "Raju", "Rakesh", "Gaurav", "Arun")

	// fmt.Println("Array:", array)
	// fmt.Println("Slice:", slice)
	// fmt.Println("Length of Slice:", len(slice))
	// fmt.Println("Capacity of Slice:", cap(slice))

	// for _, ele := range slice {                                 // _ is blank index (used when we don't care about the index)
	// 	fmt.Printf("Element=%s\n", ele)
	// }

	sort.Strings(slice)
	fmt.Println("Sorted Slice:", slice)

}
