package Validation

type AccountValidator interface {
	Validate(account string) bool
}
