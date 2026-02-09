/* Syntax:

func func_name(parameters)(return_type){
// Function body
}
*/

package main

import "fmt"

// func divide(a, b float32) float32 {
// 	if b == 0 {
// 		fmt.Println("Cannot Divide By Zero")
// 		return 0
// 	}
// 	return a / b
// }

// func main() {
// 	result := divide(1, 3)
// 	fmt.Println(result)
// }

// Call By Value:
// func multiply(a, b int) int {
// 	a = a * 2
// 	return a * b
// }

// func main() {
// 	x := 5
// 	y := 15

// 	fmt.Printf("Before Multiplication: x= %d, y= %d\n", x, y)
// 	result := multiply(x, y)
// 	fmt.Printf("After Multiplication: x= %d, y= %d\n", x, y)
// 	fmt.Printf("The Result Of Multiplication Is: %d\n", result)

// }

// Call By Reference:

func multiply(a, b *int) int {
	*a = *a * 3
	return *a * *b
}

func main() {
	x := 5
	y := 15

	fmt.Printf("Before Multiplication: x= %d, y= %d\n", x, y)
	result := multiply(&x, &y)
	fmt.Printf("After Multiplication: x= %d, y= %d\n", x, y)
	fmt.Printf("The Result Of Multiplication Is: %d\n", result)

}
