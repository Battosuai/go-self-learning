package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//Other variants of the add function
// func add(x, y int) int
func add(x int, y int) int {
	return x + y
}

// multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
// return statement without arguments returns the named return values known as "naked return"
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// If initializer is present, type can be omitted
// var c, python, java = true, false, "no!"
var c, python, java bool

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Current time is:", time.Now())
	fmt.Println("Random number:", rand.Intn(10))
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	fmt.Println("PI:", math.Pi)
	fmt.Println("Add:",add(3, 4))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	// inside a function, := is shorthand for declaring and initializing a variable
	// k := 3
	// outside a function, := construct is not available
	// Variables declared without an explicit initial value are given their zero value
	// 0 for numeric types
	// false for the boolean type
	// "" (the empty string) for strings
	var i int

	// type conversion
	f := float64(i)

	fmt.Println(i, c, python, java, f)

	// Constants cannot be declared using the := syntax
	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big))
}