package main

import "fmt"

func main() {
	// i := 0
	// for i < 4 {
	// 	i += 2
	// }
	// fmt.Println(i)

	// using in range

	// rvariable := []string{"BUBBLE", "INSERTION", "SELECTION", "MERGE"}
	// for i, j := range rvariable {
	// 	fmt.Println(i, j)
	// }

	// for strings

	// for i, j := range "abcde" {
	// 	fmt.Printf("The Index Number Of %U is %d\n", j, i)
	// }

	//for maps and key pairs

	// mmap := map[int]string{
	// 	1: "One",
	// 	2: "Two",
	// 	3: "Three",
	// 	4: "Four",
	// }

	// for key, value := range mmap {
	// 	fmt.Println(key, value)
	// }

	// For Channel:

	chnl := make(chan int)
	go func() {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()

	for i := range chnl {
		fmt.Println(i)
	}

}
