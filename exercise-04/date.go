package main

import (
	"fmt"
)

type weight struct {
	weight float64
}

func main() {
	var dob int

	fmt.Println("What is your DoB?")

	_, err := fmt.Scanf("%d", &dob)
	if err != nil{
		fmt.Printf("ERROR: invalide dob %v\n", err)
	}

	fmt.Println("What is your age?")

	var age int

	_, err = fmt.Scanf("%d", &age)
	if err != nil{
		fmt.Printf("ERROR: invalide age %v\n", err)
	}

	year := func() int{
		return dob + age
	}()
	fmt.Println("The year is", year)

	p1 := weight{
		weight: 150.0,
	}
	p2 := weight{
		weight: 89.0,
	}
	p3 := weight{
		weight: 300.0,
	}
	p4 := weight{
		weight: 512.0,
	}
	p5 := weight{
		weight: 220.0,
	}
	aw := func() float64{
		pwa := (p1.weight + p2.weight + p3.weight + p4.weight + p5.weight) / 5
		return pwa
	}()

	fmt.Printf("The average weight of the 5 provided people is : %v\n", aw)
}
