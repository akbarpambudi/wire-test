package account

import (
	"math/big"
	"wiretest/generator"
)

type InMemoryAccountGenerator struct {
	accountNumberGenerator generator.Generator
}

func (i *InMemoryAccountGenerator) GenerateAccount() (*Account, error) {
	accountNumber, err := i.accountNumberGenerator.Next()
	if err != nil {
		return nil, err
	}
	account := NewAccount(accountNumber, big.NewInt(0))
	return account, nil
}

func NewInMemoryAccountGenerator(generator generator.Generator) *InMemoryAccountGenerator {
	return &InMemoryAccountGenerator{accountNumberGenerator: generator}
}
