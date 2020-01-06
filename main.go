package main

import (
	"fmt"
	"wiretest/injector"
)

func main() {
	accountGenerator := injector.InitializeAccountGenerator()
	account, err := accountGenerator.GenerateAccount()
	if err != nil {
		panic(err)
	}
	fmt.Println(account.AccountNumber)
}
