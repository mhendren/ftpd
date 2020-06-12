package Validation

type AuthenticationType int

const (
	NotAuthenticated = iota
	Authenticated
	Anonymous
)

func (at AuthenticationType) String() string {
	return [...]string{"Not Authenticated", "Authenticated", "Anonymous"}[at]
}

type PasswordValidator interface {
	Authenticate(user string, password string) AuthenticationType
}
