package account_test

import (
	"math/big"
	"testing"
	"wiretest/account"
	"wiretest/generator"
)

func TestAccountGeneratorShouldGenerateAccount(t *testing.T) {
	generator := generator.NewAccountNumberGenerator("123", 6)
	accountGenerator := account.NewInMemoryAccountGenerator(generator)
	account, err := accountGenerator.GenerateAccount()
	if err != nil {
		t.Error(err)
	}
	if account.AccountNumber != "123000001" || account.Balance.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("account number expected to be 123000001 and balance expected to be 0, but actual account number was %s and actual balance was %s", account.AccountNumber, account.Balance.String())
	}
}
