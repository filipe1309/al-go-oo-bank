package main

import "fmt"

type CurrentAccount struct {
	holder     string
	numAgency  int
	numAccount int
	balance    float64
}

func main() {
	fmt.Println(CurrentAccount{})

	// "Filipe"
	// 589
	// 123465
	// 125.50

	// "Bob"
	// 222
	// 22222
	// 222.22

}
