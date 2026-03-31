package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Hello, To, The, Language, Go@@"
	fmt.Println("The Original String Is: ", str)

	result := strings.SplitAfterN(str, ",", 3)
	fmt.Println("The Split String Is: ", result)
}
