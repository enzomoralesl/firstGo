package main

import "fmt"

// & return variable address
// * points to address (to show or manipulate the data inside it)
//if you print only the pointer without *, you will see the address it points to

func main() {
	a, b := 33, 2701

	p := &a         // point to a
	fmt.Println(*p) // read a through the pointer
	*p = 77         // set a through the pointer
	fmt.Println(a)  // see the new value of a

	fmt.Println("a address ->", p)

	p = &b          // point to b
	fmt.Println(*p) // read b through the pointer
	*p = *p / 37    // divide b through the pointer
	fmt.Println(b)  // see the new value of b

	fmt.Println("b address ->", p)
}
