package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %s", err)
	}



	fmt.Println("Your name is", name)
	fmt.Println("Please enter a number")
	var num int
	_, err = fmt.Scanf("%d", &num)
	if err != nil{
		fmt.Println("Error number is empty", err)
	}
	if num <= 10 {
		fmt.Printf("%v is less than or equal to 10\n", num)
	}else{
		fmt.Printf("%v is greater than 10\n", num)
	}
}