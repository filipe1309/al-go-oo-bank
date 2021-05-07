package main

import (
	"accounts"
	"fmt"
)

func main() {
	johnAccount := accounts.CurrentAccount{Holder: "John", Balance: 300}
	bobAccount := accounts.CurrentAccount{Holder: "Bob", Balance: 100}

	status := johnAccount.Transfer(-200, &bobAccount)

	fmt.Println(status)
	fmt.Println(johnAccount)
	fmt.Println(bobAccount)
}
