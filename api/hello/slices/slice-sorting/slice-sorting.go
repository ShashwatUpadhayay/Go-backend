package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"Redis", "MongoDB", "Cassandra", "DynamoDB", "CouchDB"}

	fmt.Println("Before Sorting:", slice)
	fmt.Printf("Are Strings Sorted? %v\n", sort.StringsAreSorted(slice))

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})

	fmt.Println("After Sorting:", slice)
	fmt.Printf("Are Strings Sorted? %v\n", sort.StringsAreSorted(slice))
}
