package generator

import "fmt"

type AccountNumberGenerator struct {
	prefix  string
	padding int
	counter int
}

func (ag *AccountNumberGenerator) Next() (string, error) {
	ag.increaseCounter()
	paddingFormat := fmt.Sprint(ag.prefix, "%0", ag.padding, "d")
	return fmt.Sprintf(paddingFormat, ag.counter), nil
}

func (ag *AccountNumberGenerator) increaseCounter() {
	ag.counter++
}

type AccountNumberGeneratorConfig struct {
	Prefix  string
	Padding int
}

func NewAccountNumberGenerator(config AccountNumberGeneratorConfig) *AccountNumberGenerator {
	return &AccountNumberGenerator{prefix: config.Prefix, padding: config.Padding, counter: 0}
}
