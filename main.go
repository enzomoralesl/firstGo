package main

import "fmt"

func main() {

	fmt.Println("Luffy is Joyboy since")

	fmt.Println(sum(0))

	fmt.Println("years ago...")
}

func sum(sum int) int {
	c := 1
	for ; c <= 800; c++ {
		sum += 1
	}
	return sum
}
