# DevOnTheRun Notes

> notes taken during the course

```sh
go run main.go
```

```go
filipesAccount := CurrentAccount{holder: "Filipe", balance: 125.50}
```

https://play.golang.org/p/kVcneO6nVdS
https://play.golang.org/p/gkwKo7rholt

```go
filipesAccount := CurrentAccount{holder: "Filipe", numAgency: 589, numAccount: 123456, balance: 125.50}
bobsAccount := CurrentAccount{"Bob", 222, 22222, 222.22}
fmt.Println(filipesAccount, bobsAccount)

var johnAccount *CurrentAccount
johnAccount = new(CurrentAccount)
johnAccount.holder = "John"
johnAccount.balance = 500

var johnAccount2 *CurrentAccount
johnAccount2 = new(CurrentAccount)
johnAccount2.holder = "John"
johnAccount2.balance = 500

fmt.Println(*johnAccount)
fmt.Println(johnAccount, johnAccount2)
fmt.Println(johnAccount == johnAccount2)   // Comparing address -> !=
fmt.Println(*johnAccount == *johnAccount2) // Comparing values -> ==


johnAccount.holder = "John"
johnAccount.balance = 500

fmt.Println(johnAccount.balance)
status, value := johnAccount.Deposit(2000)
fmt.Println(status, value)
fmt.Println(johnAccount.Withdraw(300))
fmt.Println(johnAccount.balance)
```

# Class 3

https://blog.golang.org/using-go-modules

```sh
# To use Go Modules and locad local dependencies
go mod init bank
```

# Class 4

```go
johnAccount := accounts.CurrentAccount{Holder: "John", Balance: 300}
bobAccount := accounts.CurrentAccount{Holder: "Bob", Balance: 100}

status := johnAccount.Transfer(-200, &bobAccount)

fmt.Println(status)
fmt.Println(johnAccount)
fmt.Println(bobAccount)
```
