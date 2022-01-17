/*

Using the STRUCT “person”, using a composite literal create a value of type person and
assign it to a variable with the identifier “p1”;
print out “p1”;
print out just the field fName for "p1"

*/

package main

import "fmt"

type person struct {
	fname string
	lname string
}

func main() {
	p1 := person{
		"Ade",
		"Tokunbo",
	}

	fmt.Println(p1)
	fmt.Println(p1.fname)
}
