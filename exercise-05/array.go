package main

import "fmt"

func main() {

	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := []string{"Mason", "Taylor", "Mike", "David", "Smaug", "Bilbo", "Ophelia", "Raiden"}

	fmt.Println(s1, s2)
	for _, value := range s1 {
		fmt.Println(value)
	}
	for _, value := range s2 {
		fmt.Println(value)
	}
}
