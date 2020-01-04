package generator_test

import (
	"testing"
	"wiretest/generator"
)

func TestGeneratorShouldValueWithRightFormat(t *testing.T) {
	config := generator.AccountNumberGeneratorConfig{
		Prefix:  "1101",
		Padding: 6,
	}
	accountNumberGenerator := generator.NewAccountNumberGenerator(config)
	accountNumber, err := accountNumberGenerator.Next()
	if err != nil {
		t.Error(err)
	}
	if accountNumber != "1101000001" {
		t.Errorf("accountNumber expected to be 1101000001 but it was %s", accountNumber)
	}
}

func TestGeneratorIncreaseGeneratedValue(t *testing.T) {
	config := generator.AccountNumberGeneratorConfig{
		Prefix:  "1101",
		Padding: 6,
	}
	accountNumberGenerator := generator.NewAccountNumberGenerator(config)
	accountNumber, err := accountNumberGenerator.Next()
	if err != nil {
		t.Error(err)
	}
	accountNumber, err = accountNumberGenerator.Next()
	if err != nil {
		t.Error(err)
	}
	if accountNumber != "1101000002" {
		t.Errorf("accountNumber expected to be 1101000002 but it was %s", accountNumber)
	}
}
