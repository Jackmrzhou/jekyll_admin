package auth

import "errors"

type TokenAuthenticator struct {
	token string
}

func (ta *TokenAuthenticator) Authenticate(input interface{}) (bool, error) {
	if i, ok := input.(string); ok{
		return i == ta.token, nil
	}
	return false, errors.New("Type assertion failed: input must be a string ")
}

func NewTokenAuthenticator(token string) *TokenAuthenticator {
	return &TokenAuthenticator{token:token}
}