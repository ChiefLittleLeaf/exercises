package main

import "fmt"

func main() {
	h := greeting()
	a := address()
	fmt.Println(h)
	fmt.Println(a)
}

/* greeting returns a greeting plus my name.*/
func greeting() string {
	/* Returns string*/
	return "Hello World! I am Mason Shane Haddock!"
}

// address returns the address that i live at
func address() string {
	// returns string
	return "I live at 44 Kachina Trail, Noel Missouri"
}

