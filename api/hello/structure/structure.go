package main

import "fmt"

type Details struct {
	name       string
	age        int
	rollno     string
	gender     string
	fatherName string
}

func main() {
	var a Details
	fmt.Println(a)

	a1 := Details{"Bob", 20, "A12XB", "M", "James"}
	fmt.Println("Details1:", a1)

	fmt.Println("Name:", a1.name)
}
