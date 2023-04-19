package encrypt

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func BcryptE(ctx string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(ctx), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt error: ", err)
		return ""
	}
	return string(hash)
}

func BcryptV(hash, ctx string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(ctx))
	if err != nil {
		return false
	}
	return true
}
