package main

import "fmt"

type Author struct {
	name   string
	age    int
	salary int
}

type Teacher struct {
	details Author
}

func main() {
	result := Teacher{
		details: Author{"Bob", 30, 50000},
	}
	fmt.Println(result)
}
