package main

import (
	"fmt"
	"wiretest/injector"
)

func main() {
	accountGenerator, err := injector.InitializeAccountGenerator()
	if err != nil {
		panic(err)
	}
	account, err := accountGenerator.GenerateAccount()
	fmt.Println(account.AccountNumber)
}
