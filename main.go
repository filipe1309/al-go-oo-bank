package main

import (
	"accounts"
	"customers"
	"fmt"
)

func main() {
	johnCustomer := customers.Holder{"John", "132", "Dev"}
	johnAccount := accounts.CurrentAccount{Holder: johnCustomer, NumAgency: 123, NumAccount: 123456, Balance: 100}

	fmt.Println(johnAccount)
}
