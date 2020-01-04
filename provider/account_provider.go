package provider

import (
	"wiretest/account"
	"wiretest/generator"
)

func ProvideAccountGenerator(generator generator.Generator) account.AccountGenerator {
	return account.NewInMemoryAccountGenerator(generator)
}
