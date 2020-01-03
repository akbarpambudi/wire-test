package account

type AccountGenerator interface {
	GenerateAccount() (*Account, error)
}
