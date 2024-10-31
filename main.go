package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// main function
func main() {
	fmt.Println("Luffy is Joyboy since")

	fmt.Println(sum(0))

	fmt.Println("years ago...")

	fmt.Println(
		pow(3, 2, 9.1),
		pow(3, 3, 25.8),
	)

	operatingSystemType()
	fridayFromNow()
	timeNow()

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i + 1)
	}

	fmt.Println("done")

	//	fmt.Println(Sqrt(2))

}

//functions (func)

// looping
func sum(sum int) int {
	c := 1
	for ; c <= 800; c++ {
		sum += 1
	}
	return sum
}

// if with local variable (else can use value too)
func pow(x, y, limit float64) float64 {
	if value := math.Pow(x, y); value < limit {
		return value
	}
	return limit
}

// Newton math for approximate square root
func Sqrt(x float64) float64 {
	var z float64 = 1
	for c := 1; c <= 7; c++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%v: %v\n", c, z)
	}
	return z
}

// OS type
func operatingSystemType() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Windows.")
	default:
		// freebsd, openbsd, plan9...
		fmt.Printf("%s!\n", os)
	}
}

// days
func fridayFromNow() {
	fmt.Println("When's Friday?")
	today := time.Now().Weekday()
	switch time.Friday {
	case today:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// time
func timeNow() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
