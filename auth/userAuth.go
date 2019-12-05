package auth

import "errors"

type UserAuth struct {
	Username string
	Password string
}

type UserAuthenticator struct {
	username string
	password string
}

func (ua *UserAuthenticator) Authenticate(input interface{}) (bool, error) {
	if u, ok := input.(*UserAuth); ok {
		return u.Username == ua.username && u.Password == ua.password, nil
	}
	return false, errors.New("Type assertion failed: input type must be UserAuth type ")
}

func NewUserAuthenticator(username, password string) *UserAuthenticator {
	return &UserAuthenticator{
		username: username,
		password: password,
	}
}