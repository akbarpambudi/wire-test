package account_test

import (
	"math/big"
	"testing"
	"wiretest/account"
	"wiretest/generator"

	"github.com/stretchr/testify/assert"
)

func TestAccountGeneratorShouldGenerateAccount(t *testing.T) {
	expectedAccountNumber := "12300001"
	expectedBalance := big.NewInt(0)
	generator := generator.NewMockGenerator()
	generator.On("Next").Return(expectedAccountNumber, nil)
	accountGenerator := account.NewInMemoryAccountGenerator(generator)

	account, err := accountGenerator.GenerateAccount()

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, account.AccountNumber, expectedAccountNumber)
	assert.Equal(t, account.Balance, expectedBalance)
}
