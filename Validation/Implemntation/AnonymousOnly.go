package Implemntation

import (
	"FTPserver/Validation"
	"strings"
)

type AnonymousOnlyValidator struct {
	AllowAnonymous bool
	AnonymousUsers []string
}

func (aov AnonymousOnlyValidator) Authenticate(user string, _ string) Validation.AuthenticationType {
	isAnon := func(userName string) bool {
		lcUserName := strings.ToLower(userName)
		for _, anon := range aov.AnonymousUsers {
			if strings.ToLower(anon) == lcUserName {
				return true
			}
		}
		return false
	}
	if !aov.AllowAnonymous {
		return Validation.AuthenticationType(Validation.NotAuthenticated)
	}
	if isAnon(user) {
		return Validation.AuthenticationType(Validation.Anonymous)
	}
	return Validation.AuthenticationType(Validation.NotAuthenticated)
}
