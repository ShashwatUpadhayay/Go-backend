package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "Hello"
	s2 := "GO"
	s3 := "Lang"

	fmt.Println("strings.Compare(s1 , s2):", strings.Compare(s1, s2))
	fmt.Println("strings.Compare(s1 , s3):", strings.Compare(s1, s3))
	fmt.Println("strings.Compare(s2 , s3):", strings.Compare(s2, s3))
}
