package main

import "fmt"

// func main() {
// 	var x int16 = 432
// 	var y int16 = 40

// 	fmt.Printf("Addition: %d + %d = %d\n", x, y, x+y)
// 	fmt.Printf("Addition: %d - %d = %d\n", x, y, x-y)
// 	fmt.Printf("Addition: %d * %d = %d\n", x, y, x*y)
// 	fmt.Printf("Addition: %d / %d = %d\n", x, y, x/y)
// 	fmt.Printf("Addition: %d %% %d = %d\n", x, y, x%y)

// 	fmt.Printf("The Type Of X is %T\n", x)
// }

func main() {

	var a complex128 = complex(8, 5)
	var b complex128 = complex(6, 3)

	fmt.Printf("Addition: %v + %v = %v\n", a, b, a+b)

	fmt.Printf("The Type Of a is %T and"+"The Type Of b is %T\n", a, b)
}
