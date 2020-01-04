package provider

import "wiretest/generator"

func ProvideAccountNumberGeneratorConfig() generator.AccountNumberGeneratorConfig {
	return generator.AccountNumberGeneratorConfig{
		Prefix:  "123",
		Padding: 6,
	}
}

func ProvideAccountNumberGenerator(config generator.AccountNumberGeneratorConfig) generator.Generator {
	return generator.NewAccountNumberGenerator(config)
}
