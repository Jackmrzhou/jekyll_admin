package auth

import "jekyll_admin/conf"

type Authenticator interface {
	Authenticate(input interface{}) (bool, error)
}

func CreateAuthenticator(config *conf.Config) Authenticator {
	if config.Token != "" {
		return NewTokenAuthenticator(config.Token)
	}
	return NewUserAuthenticator(config.User.Username, config.User.Password)
}