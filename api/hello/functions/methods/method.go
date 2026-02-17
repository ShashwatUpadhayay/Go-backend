// Go methods are like functions but with a key difference: they have a receiver argument, which allows access to the receiver's properties.
// The receiver can be a struct or non-struct type, but both must be in the same package.
//  Methods cannot be created for types defined in other packages, including built-in types like int or string;
//  otherwise, the compiler will raise an error.

package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) greet() {
	fmt.Println("Name", p.name)
	fmt.Println("Age", p.age)
}

func (p *person) changeName(newName string, newAge int) {
	p.name = newName
	p.age = newAge
}
func main() {
	a := person{name: "Alice", age: 25}
	a.greet()

	fmt.Println("Before:", a.name, a.age)
	a.changeName("Alicia Jones", 18)

	fmt.Println("After:\n", a.name, a.age)

}
