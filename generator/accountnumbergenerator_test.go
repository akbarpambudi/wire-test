package generator_test

import (
	"testing"
	"wiretest/generator"

	"github.com/stretchr/testify/assert"
)

func TestGeneratorShouldValueWithRightFormat(t *testing.T) {
	expectedAccountNumber := "1101000001"
	config := generator.AccountNumberGeneratorConfig{
		Prefix:  "1101",
		Padding: 6,
	}
	accountNumberGenerator := generator.NewAccountNumberGenerator(config)

	accountNumber, err := accountNumberGenerator.Next()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, accountNumber, expectedAccountNumber)
}

func TestGeneratorIncreaseGeneratedValue(t *testing.T) {
	expectedAccountNumber := "1101000002"
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

	assert.Equal(t, expectedAccountNumber, accountNumber)

}
