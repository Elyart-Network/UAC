package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func BcryptE(ctx string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(ctx), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func BcryptV(hash, ctx string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(ctx))
	if err != nil {
		return false, err
	}
	return true, nil
}
